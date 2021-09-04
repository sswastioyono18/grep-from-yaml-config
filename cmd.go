package main

import (
	"github.com/spf13/cobra"
	"os"
)

// consumerCmd represents the version command
var consumerCmd = &cobra.Command{
	Use:   "yaml",
	Short: "grep from yaml",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := os.Getwd()
		GrepFromFile(&YamlContent{
			directory: dir,
		})
	},
}

