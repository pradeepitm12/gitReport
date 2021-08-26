package test

import (
	"bytes"
	"github.com/spf13/cobra"
)

func ExecuteCommand(root *cobra.Command, args ...string) (string, error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetArgs(args)
	_, err := root.ExecuteC()
	return buf.String(), err
}
