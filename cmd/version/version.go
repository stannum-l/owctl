package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Version struct {
	Version   string `json:"version,omitempty" yaml:"version,omitempty"`
	BuildDate string `json:"buildDate,omitempty" yaml:"buildDate,omitempty"`
	GoVersion string `json:"goVersion,omitempty" yaml:"goVersion,omitempty"`
	OS        string `json:"os,omitempty" yaml:"os,omitempty"`
	Arch      string `json:"arch,omitempty" yaml:"arch,omitempty"`
	Commit    string `json:"commit,omitempty" yaml:"commit,omitempty"`
}

func NewVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Get the version of owctl",
		Long:  "Get the version of owctl",
		Args:  cobra.NoArgs,
		RunE:  version,
	}
	return cmd
}

func version(_ *cobra.Command, _ []string) error {
	fmt.Printf("os")
	return nil
}
