package gcp

import (
	"fmt"
	"github.com/spf13/cobra"
)

func DeleteCloudRunCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "cr",
		Short:   "Delete an CloudRun container",
		Long:    `Delete an CloudRun container`,
		GroupID: "gcp",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("deleting CloudRun container")
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&name, "name", "n", "", "Name of the AKS cluster")

	return cmd
}
