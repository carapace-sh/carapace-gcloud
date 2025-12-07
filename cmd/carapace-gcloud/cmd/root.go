package cmd

import (
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-gcloud/cmd/carapace-gcloud/cmd/gcloud"
	spec "github.com/carapace-sh/carapace-spec"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "gcloud",
	Short: "An enriched gloud completer",
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	rootCmd.SetUsageFunc(func(c *cobra.Command) error { return nil })
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("access-token-file", "", "A file path to read the access token.")
	rootCmd.PersistentFlags().String("account", "", "Google Cloud user account to use for invocation.")
	rootCmd.PersistentFlags().String("authority-selector", "", "THIS ARGUMENT NEEDS HELP TEXT.")
	rootCmd.PersistentFlags().String("authorization-token-file", "", "THIS ARGUMENT NEEDS HELP TEXT.")
	rootCmd.PersistentFlags().String("billing-project", "", "The Google Cloud project that will be charged quota for")
	rootCmd.PersistentFlags().String("configuration", "", "File name of the configuration to use for this command invocation.")
	rootCmd.PersistentFlags().String("credential-file-override", "", "THIS ARGUMENT NEEDS HELP TEXT.")
	rootCmd.PersistentFlags().String("document", "", "THIS TEXT SHOULD BE HIDDEN")
	rootCmd.PersistentFlags().String("flags-file", "", "A YAML or JSON file that specifies a *--flag*:*value* dictionary.")
	rootCmd.PersistentFlags().String("flatten", "", "Flatten _name_[] output resource slices in _KEY_ into separate records")
	rootCmd.PersistentFlags().String("force-endpoint-mode", "", "Regional endpoint mode to use.")
	rootCmd.PersistentFlags().String("format", "", "Sets the format for printing command output resources.")
	rootCmd.PersistentFlags().Bool("help", false, "Display detailed help.")
	rootCmd.PersistentFlags().String("http-timeout", "", "THIS ARGUMENT NEEDS HELP TEXT.")
	rootCmd.PersistentFlags().String("impersonate-service-account", "", "For this `gcloud` invocation, all API requests will be")
	rootCmd.PersistentFlags().Bool("log-http", false, "Log all HTTP server requests and responses to stderr.")
	rootCmd.PersistentFlags().Bool("no-log-http", false, "Log all HTTP server requests and responses to stderr.")
	rootCmd.PersistentFlags().Bool("no-user-output-enabled", false, "Print user intended output to the console.")
	rootCmd.PersistentFlags().String("project", "", "The Google Cloud project ID to use for this invocation.")
	rootCmd.PersistentFlags().Bool("quiet", false, "Disable all interactive prompts when running `gcloud` commands.")
	rootCmd.PersistentFlags().String("trace-token", "", "Token used to route traces of service requests for investigation of issues.")
	rootCmd.PersistentFlags().String("universe-domain", "", "Universe domain to target.")
	rootCmd.PersistentFlags().Bool("user-output-enabled", false, "Print user intended output to the console.")
	rootCmd.PersistentFlags().String("verbosity", "", "Override the default verbosity for this command.")
	rootCmd.PersistentFlags().Bool("version", false, "Print version information and exit.")
	rootCmd.Flag("authority-selector").Hidden = true
	rootCmd.Flag("authorization-token-file").Hidden = true
	rootCmd.Flag("credential-file-override").Hidden = true
	rootCmd.Flag("document").Hidden = true
	rootCmd.Flag("force-endpoint-mode").Hidden = true
	rootCmd.Flag("http-timeout").Hidden = true
	rootCmd.Flag("no-log-http").Hidden = true
	rootCmd.Flag("no-user-output-enabled").Hidden = true
	rootCmd.Flag("universe-domain").Hidden = true

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"access-token-file":           carapace.ActionValues(),
		"account":                     carapace.ActionValues(),
		"authority-selector":          carapace.ActionValues(),
		"authorization-token-file":    carapace.ActionValues(),
		"billing-project":             carapace.ActionValues(),
		"configuration":               carapace.ActionValues(),
		"credential-file-override":    carapace.ActionValues(),
		"document":                    carapace.ActionValues(),
		"flags-file":                  carapace.ActionValues(),
		"flatten":                     carapace.ActionValues(),
		"force-endpoint-mode":         carapace.ActionValues(),
		"format":                      carapace.ActionValues(),
		"http-timeout":                carapace.ActionValues(),
		"impersonate-service-account": carapace.ActionValues(),
		"project":                     carapace.ActionValues(),
		"trace-token":                 carapace.ActionValues(),
		"universe-domain":             carapace.ActionValues(),
		"verbosity":                   carapace.ActionValues(),
	})

	for name, description := range gcloud.Services() {
		serviceCmd := &cobra.Command{
			Use:   name,
			Short: description,
			Run:   func(cmd *cobra.Command, args []string) {},
		}
		carapace.Gen(serviceCmd).Standalone()
		rootCmd.AddCommand(serviceCmd)
		carapace.Gen(serviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
			gcloudCommand, err := gcloud.Get(fmt.Sprintf("gcloud.%s.yaml", serviceCmd.Use))
			if err != nil {
				carapace.LOG.Println(err.Error()) // TODO handle error
				return
			}

			for _, subCmd := range gcloudCommand.Commands {
				operationCmd := spec.Command(subCmd).ToCobra()
				serviceCmd.AddCommand(operationCmd)

				carapace.Gen(operationCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
					// TODO same for deeper subcommands
					if flag != nil && flag.Value.Type() != "bool" {
						if _, ok := subCmd.Completion.Flag[flag.Name]; !ok {
							return carapace.ActionMessage("TODO bridge gcloud completer")
						}
					}
					return action
				})
			}
		})
	}

	spec.Register(rootCmd)
}
