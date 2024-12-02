package gcp

import (
	"context"
	"fmt"
	"time"

	container "cloud.google.com/go/container/apiv1"
	"cloud.google.com/go/container/apiv1/containerpb"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

type GKE struct {
	Name    string `json:"name,omitempty" yaml:"name,omitempty"`
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
	Region  string `json:"region,omitempty" yaml:"region,omitempty"`
}

var cluster GKE

func NewGKECommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "gcp",
		Short:   "Create a GKE cluster",
		Long:    `Create a GKE cluster`,
		GroupID: "gcp",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("starting GKE cluster is not yet implemented")
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&cluster.Name, "name", "n", "", "Name of the AKS cluster")
	f.StringVarP(&cluster.Version, "version", "v", "1.29.0", "Kubernetes Version of the AKS cluster")
	f.StringVarP(&cluster.Region, "region", "r", "", "AWS Region")

	cobra.MarkFlagRequired(f, "name")
	cobra.MarkFlagRequired(f, "region")
	return cmd
}

func createGKECluster(ctx context.Context, projectID, zone, network, subnet, clusterName, version string) error {

	gkeClient, err := container.NewClusterManagerClient(ctx, option.WithEndpoint("https://container.googleapis.com"))
	if err != nil {
		return fmt.Errorf("failed to create GKE client: %v", err)
	}
	defer gkeClient.Close()

	cluster := &containerpb.Cluster{
		Name:       clusterName,
		Network:    network,
		Subnetwork: subnet,
	}

	req := &containerpb.CreateClusterRequest{
		ProjectId: projectID,
		Zone:      zone,
		Cluster:   cluster,
	}

	op, err := gkeClient.CreateCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to create GKE cluster: %v", err)
	}

	// wait
	for {
		op, err = gkeClient.GetOperation(ctx, &containerpb.GetOperationRequest{
			ProjectId:   projectID,
			Zone:        zone,
			OperationId: op.Name,
		})

		if err != nil {
			return fmt.Errorf("failed to get operation: %v", err)
		}

		if op.Status == containerpb.Operation_DONE {
			if op.StatusMessage != "" {
				fmt.Printf("Operation completed successfully: %s\n", op.StatusMessage)
			}
			fmt.Printf("Waiting for GKE cluster %s to become ready...\n", clusterName)
			break
		}

		time.Sleep(30 * time.Second)
	}
	fmt.Printf("GKE cluster %s created successfully\n", clusterName)
	return nil

}
