package create

import (
	"github.com/spf13/cobra"

	"github.com/stannum-l/owctl/cmd/create/aws"
	"github.com/stannum-l/owctl/cmd/create/azure"
	"github.com/stannum-l/owctl/cmd/create/gcp"
)

func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create cloud native container and cluster resources",
		Long:  `Create cloud native container and cluster resources`,
	}

	cmd.AddGroup(
		&cobra.Group{Title: "Microsoft Azure:", ID: "azure"},
		&cobra.Group{Title: "Amazon Web Services (AWS):", ID: "aws"},
		&cobra.Group{Title: "Google Cloud Platform (GCP):", ID: "gcp"},
	)

	cmd.AddCommand(azure.NewAKSCommand())
	cmd.AddCommand(aws.NewEKSCommand())
	cmd.AddCommand(gcp.NewGKECommand())

	//cmd.AddCommand(azure.NewACACommand())
	//cmd.AddCommand(aws.NewECSCommand())
	//cmd.AddCommand(gcp.NewCloudRunCommand())

	return cmd
}
