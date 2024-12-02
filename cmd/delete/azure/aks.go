package azure

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
	"github.com/spf13/cobra"
)

var name string

func DeleteAKSCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "aks",
		Short:   "Delete an Azure cluster.",
		Long:    `Delete an Azure cluster.`,
		GroupID: "azure",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("deleting AKS cluster")
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&name, "name", "n", "", "Name of the AKS cluster")

	return cmd
}

func deleteCluster(clusterName, subscriptionID, resourceGroupName string) error {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return fmt.Errorf("failed to obtain a credential: %v", err)
	}

	ctx := context.Background()
	client, err := armcontainerservice.NewManagedClustersClient(subscriptionID, cred, nil)
	if err != nil {
		return fmt.Errorf("failed to obtain a container %v", err)
	}

	poller, err := client.BeginDelete(ctx, resourceGroupName, clusterName, nil)
	if err != nil {
		return fmt.Errorf("failed to begin deleting cluster: %v", err)
	}

	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to poll deleting cluster: %v", err)
	}

	fmt.Printf("successfully deleting AKS cluster %s\n", clusterName)
	return nil
}
