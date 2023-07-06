package cmd

import (
	"fmt"
	"github.com/carlmjohnson/versioninfo"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "goreleaser-testing",
	Short:   "Testing goreleaser",
	Long:    "Testing goreleaser long",
	Run:     mainEntrypoint,
	Version: versioninfo.Short(),
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func mainEntrypoint(_ *cobra.Command, _ []string) {
	fmt.Printf("Version: %s, %s, %s\n", versioninfo.Version, versioninfo.Revision, versioninfo.Short())
}
