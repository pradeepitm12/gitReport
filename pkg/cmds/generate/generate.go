package generate

import (
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate report",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := somework(); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func somework() error {
	return nil
}
