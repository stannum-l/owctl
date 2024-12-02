package gcp

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCloudRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cloudrun",
		Short:   "Create a CloudRun container",
		Long:    `Create a CloudRun container`,
		GroupID: "gcp",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("starting Cloud Run container is not yet implemented")
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
