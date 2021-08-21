package cmd

import (
	"github.com/sika-training-examples/go-cli-example/cmd/root"
	_ "github.com/sika-training-examples/go-cli-example/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
