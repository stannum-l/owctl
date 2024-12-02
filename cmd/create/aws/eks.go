package aws

import (
	"fmt"
	"github.com/stannum-l/owctl/pkg/model"

	"github.com/spf13/cobra"
)

var cluster model.EKS

func NewEKSCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "eks",
		Short:   "Create an Elastic Kubernetes Service cluster",
		Long:    `Create an Elastic Kubernetes Service cluster`,
		GroupID: "aws",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("starting EKS cluster")
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&cluster.Name, "name", "n", "", "Name of the EKS cluster")
	f.StringVarP(&cluster.Version, "version", "v", "1.29.0", "Kubernetes Version of the AKS cluster")
	f.StringVarP(&cluster.Region, "region", "r", "", "AWS Region")

	cobra.MarkFlagRequired(f, "name")
	cobra.MarkFlagRequired(f, "region")

	return cmd
}

//func createEKSCluster(ctx context.Context, region, clusterName, subnetID, vpcID string) error {
//	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
//	if err != nil {
//		return fmt.Errorf("loading config: %w", err)
//	}
//
//	eksClient := aws.NewFromConfig(cfg)
//	eksInput := &aws.CreateClusterInput{
//		Name: aws.String(clusterName),
//	}
//
//	output, err := eksClient.CreateCluster(ctx, eksInput)
//	if err != nil {
//		return fmt.Errorf("failed to create EKS cluster: %w", err)
//	}
//
//	return nil
//}
