package delete

import (
	"github.com/spf13/cobra"
	"github.com/stannum-l/owctl/cmd/delete/aws"
	"github.com/stannum-l/owctl/cmd/delete/azure"
	"github.com/stannum-l/owctl/cmd/delete/gcp"
)

func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete cloud native container and cluster resources",
		Long:  `Delete cloud native container and cluster resources`,
	}

	cmd.AddGroup(
		&cobra.Group{Title: "Microsoft Azure:", ID: "azure"},
		&cobra.Group{Title: "Amazon Web Services (AWS):", ID: "aws"},
		&cobra.Group{Title: "Google Cloud Platform (GCP):", ID: "gcp"},
	)

	cmd.AddCommand(azure.DeleteAKSCommand())
	cmd.AddCommand(aws.DeleteEKSCommand())
	cmd.AddCommand(gcp.DeleteGKECommand())

	//cmd.AddCommand(azure.DeleteACACommand())
	//cmd.AddCommand(aws.DeleteECSCommand())
	//cmd.AddCommand(gcp.DeleteCloudRunCommand())

	return cmd
}
