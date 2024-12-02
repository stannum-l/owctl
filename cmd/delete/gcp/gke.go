package gcp

import (
	"cloud.google.com/go/container/apiv1/containerpb"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"time"

	container "cloud.google.com/go/container/apiv1"
)

var name string

func DeleteGKECommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "gke",
		Short:   "Delete an GKE cluster.",
		Long:    `Delete an GKE cluster.`,
		GroupID: "gcp",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("deleting GKE cluster")
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&name, "name", "n", "", "Name of the AKS cluster")

	return cmd
}

func deleteGKECluster(ctx context.Context, projectID, location, clusterName string) error {

	client, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to obtain a container %v", err)
	}
	defer client.Close()

	clusterPath := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", projectID, location, clusterName)
	op, err := client.DeleteCluster(ctx, &containerpb.DeleteClusterRequest{
		Name: clusterPath,
	})
	if err != nil {
		return fmt.Errorf("failed to delete GKE cluster %v", err)
	}

	for {
		status := op.GetStatus()
		if status == containerpb.Operation_DONE {
			fmt.Printf("successfully deleted GKE cluster %s\n", clusterName)
			return nil
		} else if status == containerpb.Operation_ABORTING || status == containerpb.Operation_STATUS_UNSPECIFIED {
			return fmt.Errorf("GKE cluster %s is in an unknown state: %v", clusterName, op)
		}
		time.Sleep(20 * time.Second)
		fmt.Print(".")
	}
}
