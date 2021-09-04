package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sswastioyono18/grep-from-yaml-config/app"
)

var yamlCommand = &cobra.Command{
	Use:   "yaml",
	Short: "grep from yaml and check if key is used or not in project",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		app.StartGrepFromFile(NewYamlContent())
	},
}

func init() {
	rootCmd.AddCommand(yamlCommand)
}

func NewYamlContent() *app.YamlContent {
	return &app.YamlContent{
		FileTarget: viper.GetString("FILE_TARGET"),
		TracePath: viper.GetStringSlice("YAML_TRACE_PATH"),
		Project: app.Project{
			GrepDirSource: viper.GetString("GREP_DIR_SOURCE"),
			GrepDirTarget: viper.GetString("GREP_DIR_TARGET"),
			AppTarget:     viper.GetStringSlice("APP_TARGET"),
			RootPath:      viper.GetString("ROOT_PATH"),
		},
	}
}