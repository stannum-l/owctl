package azure

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewACACommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "aca",
		Short:   "Create Azure Container App container",
		Long:    `Create Azure Container App container`,
		GroupID: "azure",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("starting Azure Container App is not yet implemented")
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&cluster.Name, "name", "n", "", "Azure Container App name")
	f.StringVarP(&cluster.Location, "location", "l", "", "Azure Container App location")

	_ = cobra.MarkFlagRequired(f, "name")

	return cmd
}
