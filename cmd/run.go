package cmd

import (
	"log"
	"os"

	"github.com/rishabh-j-23/xofikit/pkg/menu"
	"github.com/rishabh-j-23/xofikit/pkg/runner"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [config.yaml]",
	Short: "Run xofikit with specified YAML config file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configPath := args[0]

		cfg, err := menu.LoadConfig(configPath)
		if err != nil {
			log.Println("Failed to load config:", err)
			os.Exit(1)
		}

		opt, err := menu.ShowMenu(cfg)
		if err != nil {
			log.Println("Menu error:", err)
			os.Exit(1)
		}

		if opt.Command == "" {
			log.Println("No command defined for selected option.")
			os.Exit(1)
		}

		log.Printf("Executing command: %s\n", opt.Command)
		runner.ExecuteCommand(opt.Command)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
