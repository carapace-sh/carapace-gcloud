package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-spec/pkg/command"
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

		var glabCommand Command
		if err := json.Unmarshal(content, &glabCommand); err != nil {
			return err
		}

		command := glabCommand.ToSpecCommand(cli.SerializedFlagList, true)
		if cmd.Flag("no-doc").Changed {
			stripDoc(&command)
		}

		if cmd.Flag("stdout").Changed {
			m, err := yaml.Marshal(command)
			if err != nil {
				return err
			}
			fmt.Println("# yaml-language-server: $schema=https://carapace.sh/schemas/command.json")
			fmt.Println(string(m))
			return nil
		}

		dir := cmd.Flag("target").Value.String()
		if dir == "" {
			dir, err = os.MkdirTemp("", "carapace-spec-botocore-*")
			if err != nil {
				return err
			}
		}

		for _, subCommand := range command.Commands {
			m, err := yaml.Marshal(subCommand)
			if err != nil {
				return err
			}
			m = append([]byte("# yaml-language-server: $schema=https://carapace.sh/schemas/command.json\n"), m...)
			path := path.Join(dir, fmt.Sprintf("gcloud.%s.yaml", subCommand.Name))
			println(path)
			if err := os.WriteFile(path, m, os.ModePerm); err != nil {
				return err
			}
		}

		command.Commands = nil
		m, err := yaml.Marshal(command)
		if err != nil {
			return err
		}
		path := path.Join(dir, "gcloud.yaml")
		println(path)
		if err := os.WriteFile(path, m, os.ModePerm); err != nil {
			return err
		}
		return nil

	},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().Bool("no-doc", false, "strip documentation")
	rootCmd.Flags().Bool("stdout", false, "print to stdout")
	rootCmd.Flags().String("target", "", "target directory")
	rootCmd.MarkFlagsMutuallyExclusive("stdout", "target")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}

func stripDoc(command *command.Command) {
	command.Documentation.Command = ""
	command.Documentation.Flag = nil
	for index := range command.Commands {
		stripDoc(&command.Commands[index])
	}
}
