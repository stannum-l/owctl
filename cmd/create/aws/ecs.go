package aws

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewECSCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ecs",
		Short:   "Create an Elastic Container Service container",
		Long:    `Create an Elastic Container Service container`,
		GroupID: "aws",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("starting Elastic Container Service is not yet implemented")
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
