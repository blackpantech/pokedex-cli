package cmd

import (
	"log"
	"os"

	"github.com/blackpantech/pokedex-cli/internal/api"
	"github.com/blackpantech/pokedex-cli/internal/cli"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pokedex-cli",
	Short: "A simple CLI Pokedex.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalf("No name or id was given!\n")
			os.Exit(1)
		}
		pokemon, err := api.GetPokemon(args[0])
		if err != nil {
			os.Exit(1)
		}
		cli.ClearDisplay()
		cli.DisplaySprite(pokemon.Sprites.FrontDefault)
		cli.DisplayPokedexEntry(*pokemon)
		cli.PlaySound(pokemon.Cries.Latest)
	},
}

// Execute executes the root command.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
