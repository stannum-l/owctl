package aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/spf13/cobra"
)

var name string

func DeleteEKSCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "eks",
		Short:   "Delete an EKS cluster.",
		Long:    `Delete an EKS cluster.`,
		GroupID: "aws",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("deleting EKS cluster")
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&name, "name", "n", "", "Name of the AKS cluster")

	return cmd
}

func deleteManagedNodeGrouops(ctx context.Context, client *eks.Client, clusterName string) error {
	res, err := client.ListNodegroups(ctx, &eks.ListNodegroupsInput{
		ClusterName: aws.String(clusterName),
	})
	if err != nil {
		return fmt.Errorf("failed to list node groups: %v", err)
	}

	for _, n := range res.Nodegroups {
		fmt.Printf("deleting managed nodegroup: %s\n", n)
		_, err = client.DeleteNodegroup(ctx, &eks.DeleteNodegroupInput{
			ClusterName:   aws.String(clusterName),
			NodegroupName: aws.String(n),
		})
		if err != nil {
			return fmt.Errorf("failed to delete node group %s: %v", n, err)
		}
		fmt.Printf("node group %s deleted successfully\n", n)
	}
	return nil
}

func deleteCluster(clusterName string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return fmt.Errorf("unable to load SDK config, %w", err)
	}

	client := eks.NewFromConfig(cfg)

	err = deleteManagedNodeGrouops(context.TODO(), client, clusterName)
	if err != nil {
		return fmt.Errorf("failed to delete node groups: %v", err)
	}

	_, err = client.DeleteCluster(context.TODO(), &eks.DeleteClusterInput{
		Name: aws.String(clusterName),
	})
	if err != nil {
		return fmt.Errorf("failed to delete EKS cluster: %v", err)
	}

	fmt.Printf("successfully deleting EKS cluster %s\n", clusterName)
	return nil
}
