package generate

import (
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate report",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := gitwork(cmd); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.PersistentFlags().String("organization", "", "gitrep --organization=yyyyy")
	cmd.PersistentFlags().String("auth-key", "", "gitrep --auth-key=xxxxx")
	cmd.MarkPersistentFlagRequired("organization")
	cmd.MarkPersistentFlagRequired("auth-key")

	return cmd
}
