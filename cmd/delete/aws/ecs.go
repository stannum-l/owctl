package aws

import (
	"fmt"
	"github.com/spf13/cobra"
)

func DeleteECSCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "ecs",
		Short:   "Delete an Elastic Container Service container",
		Long:    `Delete an Elastic Container Service container`,
		GroupID: "aws",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("deleting ECS container")
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&name, "name", "n", "", "Name of the ECS container")

	return cmd
}
