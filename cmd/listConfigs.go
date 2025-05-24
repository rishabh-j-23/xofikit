/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// listConfigsCmd represents the listConfigs command
var listConfigsCmd = &cobra.Command{
	Use:   "list-configs",
	Short: "List available config YAML files",
	RunE: func(cmd *cobra.Command, args []string) error {

		configDir := filepath.Join(os.Getenv("HOME"), ".config", "xofikit", "scripts_config")

		newScriptsConfigPath, err := cmd.Flags().GetString("scripts-path")

		if err != nil {
			return fmt.Errorf(`Invalid file path %s`, newScriptsConfigPath)
		} else {
			configDir = newScriptsConfigPath
		}

		files, err := os.ReadDir(configDir)
		if err != nil {
			return fmt.Errorf("failed to read config directory: %w", err)
		}

		fmt.Println("Available configs:")
		for _, f := range files {
			if !f.IsDir() && strings.HasSuffix(f.Name(), ".yaml") {
				fmt.Println(" -", f.Name())
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listConfigsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listConfigsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listConfigsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listConfigsCmd.PersistentFlags().String("scripts-path", "scripts_config", "Path for custom scripts location")

}
