package cli

import (
	"github.com/SlateLH/teakit/internal/commands"
	"github.com/SlateLH/teakit/internal/config"
	"github.com/SlateLH/teakit/internal/tui"
	"github.com/spf13/cobra"
)

var initForce bool

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize teakit in the current project",
	Long:  "The init command sets up teakit in your current project by creating necessary configuration files and directories. It prepares your project to use teakit for managing components.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if initForce {
			return commands.Init(config.Default())
		} else {
			return tui.RunInit()
		}
	},
}

func init() {
	initCmd.Flags().BoolVarP(&initForce, "yes", "y", false, "run non-interactively with default options")

	rootCmd.AddCommand(initCmd)
}
