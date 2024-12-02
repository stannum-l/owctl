package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/stannum-l/owctl/cmd/create"
	"github.com/stannum-l/owctl/cmd/delete"
	"github.com/stannum-l/owctl/cmd/version"
	"github.com/stannum-l/owctl/pkg/util"
)

var (
	debug  bool
	logger *slog.Logger
)

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	logger = util.NewLogger(debug)
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "owctl",
		Short: "OW Kubernetes CLI tool",
		Long:  `A unified cloud provider Kubernetes cluster`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	// hides the completion command
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	// adds all the subcommand
	rootCmd.AddCommand(create.NewCreateCommand())
	rootCmd.AddCommand(delete.NewDeleteCommand())
	rootCmd.AddCommand(version.NewVersionCmd())

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug logging")

	if err := rootCmd.Execute(); err != nil {
		fmt.Errorf("error executing command: %v", err)
		os.Exit(1)
	}
}
