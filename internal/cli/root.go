package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "teakit",
	Short: "Installer for charmbracelet/bubbles components",
	Long:  "teakit initializes and installs reusable charmbracelet/bubbletea components into your project.",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
