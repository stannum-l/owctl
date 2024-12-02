package azure

import (
	"context"
	"fmt"
	"github.com/stannum-l/owctl/pkg/model"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
	"github.com/spf13/cobra"
)

var cluster model.AKS
var nodepool model.AzureNodepool

func NewAKSCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "aks",
		Short:   "Create an AKS cluster",
		Long:    `Create an AKS cluster`,
		GroupID: "azure",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("starting AKS cluster")
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&cluster.Name, "name", "n", "", "Name of the AKS cluster")
	f.StringVarP(&cluster.Version, "version", "v", "1.29.0", "Kubernetes Version of the AKS cluster")
	f.StringVarP(&cluster.Location, "location", "l", "", "AKS location")
	f.StringVar(&cluster.DNSPrefix, "dns-prefix", "", "DNS prefix")
	f.StringVar(&cluster.PodCIDR, "pod-cidr", "10.244.0.0/16", "Pod CIDR")
	f.StringVar(&cluster.ServiceCIDR, "service-cidr", "10.96.0.0/12", "Service CIDR")

	f.StringVar(&nodepool.Name, "nodepool-name", "", "Node pool name")

	//_ = cobra.MarkFlagRequired(f, "name")
	//_ = cobra.MarkFlagRequired(f, "region")
	return cmd
}

func createAKSCluster(ctx context.Context, c model.AKS, n model.AzureNodepool) error {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return fmt.Errorf("authentication error: %w", err)
	}

	// create an AKS client
	aksClient, err := armcontainerservice.NewManagedClustersClient(c.SubscriptionID, cred, nil)
	if err != nil {
		return fmt.Errorf("failed to create managed clusters client: %v", err)
	}

	aksParams := armcontainerservice.ManagedCluster{
		Location: to.Ptr(c.Location),
		Properties: &armcontainerservice.ManagedClusterProperties{
			DNSPrefix: to.Ptr(c.DNSPrefix),
			NetworkProfile: &armcontainerservice.NetworkProfile{
				NetworkPlugin: to.Ptr(armcontainerservice.NetworkPluginAzure),
				NetworkPolicy: to.Ptr(armcontainerservice.NetworkPolicyCalico),
				PodCidr:       to.Ptr(c.PodCIDR),
				ServiceCidr:   to.Ptr(c.ServiceCIDR),
				OutboundType:  to.Ptr(armcontainerservice.OutboundTypeUserDefinedRouting),
			},
			EnableRBAC: to.Ptr(true),
			AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
				{
					Name:              to.Ptr(n.Name),
					MinCount:          to.Ptr(n.Min),
					MaxCount:          to.Ptr(n.Max),
					VMSize:            to.Ptr(n.VMType),
					EnableAutoScaling: to.Ptr(true),
				},
			},
		},
	}

	resp, err := aksClient.BeginCreateOrUpdate(ctx, c.ResourceGroup, c.Name, aksParams, nil)
	if err != nil {
		return fmt.Errorf("failed to create managed c: %v", err)
	}

	_, err = resp.PollUntilDone(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to create managed c: %v", err)
	}

	fmt.Printf("Created managed c %s\n", c.Name)
	return nil
}
