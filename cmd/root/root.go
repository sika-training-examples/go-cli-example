
package root

import (
	"github.com/sika-training-examples/go-cli-example/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "go-cli-example",
	Short: "go-cli-example, " + version.Version,
}
