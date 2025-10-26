package cmd

import (
	"os"

	"github.com/blackpantech/pokedex-cli/internal/api"
	"github.com/blackpantech/pokedex-cli/internal/cli"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pokedex-cli",
	Short: "A simple CLI Pokedex.",
	Run: func(cmd *cobra.Command, args []string) {
		pokemon, err := api.GetPokemon(args[0])
		if err != nil {
			os.Exit(1)
		}

		cli.DisplaySprite(pokemon.Sprites.FrontDefault)
		cli.DisplayPokedexEntry(*pokemon)
		cli.PlaySound(pokemon.Cries.Latest)
	},
}

// Execute executes the root command.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
