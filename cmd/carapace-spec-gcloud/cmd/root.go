package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var rootCmd = &cobra.Command{
	Use:   "carapace-spec-gcloud",
	Short: "",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		content, err := os.ReadFile(args[0])
		if err != nil {
			return err
		}

		var cli Cli
		if err := json.Unmarshal(content, &cli); err != nil {
			return err
		}

		var command Command
		if err := json.Unmarshal(content, &command); err != nil {
			return err
		}

		m, err := yaml.Marshal(command.ToSpecCommand(cli.SerializedFlagList, true))
		if err != nil {
			return err
		}

		fmt.Println(string(m))
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
