package cmds

import (
	"github.com/pradeepitm12/gitReport/pkg/cmds/generate"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gitrep",
		Short: "generate git user report",
		Long: `gitrep stands for git user report generator
This is used to generate some usage report,
and the skill set git user is building.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(
		generate.Command(),
	)
}
