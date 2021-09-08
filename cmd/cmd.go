package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sswastioyono18/grep-from-yaml-config/app"
	"github.com/sswastioyono18/grep-from-yaml-config/log"
)

var cleanYamlConfigCmd = &cobra.Command{
	Use:   "clean-yaml-config",
	Short: "grep from yaml and check if key is used or not in project",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		log.NewLogger()
		log.Logger.Info("Starting clean-yaml-config")
		app.StartGrepFromFile(&app.YamlContent{
			TracePath: viper.GetStringSlice("CONFIG_YAML_TRACE_PATH"),
			Project: app.Project{
				GrepDirSource: viper.GetStringSlice("GREP_CONFIG_YAML_SOURCE"),
				GrepExtTarget: viper.GetString("GREP_EXT_TARGET"),
				AppTarget:     viper.GetStringSlice("APP_TARGET"),
				RootPath:      viper.GetString("ROOT_PATH"),
			},
		})
	},
}

var cleanYamlSecretCmd = &cobra.Command{
	Use:   "clean-yaml-secret",
	Short: "grep from yaml and check if key is used or not in project",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		log.NewLogger()
		log.Logger.Info("Starting clean-yaml-secret")
		app.StartGrepFromFile(&app.YamlContent{
			TracePath: viper.GetStringSlice("SECRET_YAML_TRACE_PATH"),
			Project: app.Project{
				GrepDirSource: viper.GetStringSlice("GREP_SECRET_YAML_SOURCE"),
				GrepExtTarget: viper.GetString("GREP_EXT_TARGET"),
				AppTarget:     viper.GetStringSlice("APP_TARGET"),
				RootPath:      viper.GetString("ROOT_PATH"),
			},
		})
	},
}

func init() {
	rootCmd.AddCommand(cleanYamlConfigCmd, cleanYamlSecretCmd)
}
