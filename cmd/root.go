package cmd

import (
	"github.com/blackpantech/pokedex-cli/internal/api"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pokedex-cli",
	Short: "A simple CLI Pokedex.",
	Run: func(cmd *cobra.Command, args []string) {
		api.GetPokemon(args[0])
	},
}

// Execute executes the root command.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
