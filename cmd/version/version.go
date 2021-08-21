package version

import (
	"fmt"

	"github.com/sika-training-examples/go-cli-example/cmd/root"
	"github.com/sika-training-examples/go-cli-example/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "version",
	Short:   "Prints version",
	Aliases: []string{"v"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("%s\n", version.Version)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
