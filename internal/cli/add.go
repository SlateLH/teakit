package cli

import (
	"github.com/SlateLH/teakit/internal/commands"
	"github.com/spf13/cobra"
)

var (
	addForce    bool
	addRegistry string
)

var addCmd = &cobra.Command{
	Use:   "add <component>",
	Short: "Add a component to the project",
	Long:  "The add command fetches and adds a specified component to your teakit-managed project.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		opts := commands.AddOptions{
			Component: args[0],
			Registry:  addRegistry,
			Force:     addForce,
		}

		return commands.Add(opts)
	},
}

func init() {
	addCmd.Flags().BoolVarP(&addForce, "force", "f", false, "overwrite existing component if it already exists")
	addCmd.Flags().StringVarP(&addRegistry, "registry", "r", "", "registry alias to use (from teakit.json)")

	rootCmd.AddCommand(addCmd)
}
