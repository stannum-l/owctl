package azure

import (
	"fmt"
	"github.com/spf13/cobra"
)

func DeleteACACommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "aca",
		Short:   "Delete an Azure Container App container",
		Long:    `Delete an Azure Container App container.`,
		GroupID: "azure",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("deleting Azure Container App container")
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&name, "name", "n", "", "Name of the AKS cluster")

	return cmd
}
