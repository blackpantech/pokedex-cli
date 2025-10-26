package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/blackpantech/pokedex-cli/internal/models"
	"github.com/blackpantech/pokedex-cli/internal/utils"

	"github.com/charmbracelet/glamour"
)

func DisplaySprite(url string) error {
	cmd := exec.Command("chafa", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func DisplayPokedexEntry(pokemon models.Pokemon) {
	title := getTitle(pokemon)
	types := getTypes(pokemon)
	stats := getStats(pokemon)
	abilities := getAbilities(pokemon)
	moves := getMoves(pokemon)
	in := strings.Join(
		[]string{title, types, stats, abilities, moves},
		"\n***\n",
	)
	out, _ := glamour.Render(in, "auto")
	fmt.Print(out)
}

func getTitle(pokemon models.Pokemon) string {
	return fmt.Sprintf(
		"# %s, ID: %d",
		strings.ToUpper(pokemon.Name),
		pokemon.ID,
	)
}

func getTypes(pokemon models.Pokemon) string {
	typeSlice := make([]string, 0)
	for _, t := range pokemon.Types {
		typeSlice = append(typeSlice, utils.CapitalizeFirstLetters(t.Type.Name))
	}
	return strings.Join(typeSlice, ", ")
}

func getStats(pokemon models.Pokemon) string {
	statsSlice := make([]string, 0)
	for _, s := range pokemon.Stats {
		statLine := fmt.Sprintf("- %s: %d, effort: %d",
			utils.CapitalizeFirstLetters(s.Stat.Name), s.BaseStat, s.Effort)
		statsSlice = append(statsSlice, statLine)
	}
	return fmt.Sprintf("## Stats\n %s", strings.Join(statsSlice, "\n"))
}

func getAbilities(pokemon models.Pokemon) string {
	abilitiesSlice := make([]string, 0)
	for _, a := range pokemon.Abilities {
		abilityLine := fmt.Sprintf(
			"- %s",
			utils.CapitalizeFirstLetters(a.Ability.Name),
		)
		abilitiesSlice = append(abilitiesSlice, abilityLine)
	}
	return fmt.Sprintf("## Abilities\n %s", strings.Join(abilitiesSlice, "\n"))
}

func getMoves(pokemon models.Pokemon) string {
	movesSlice := make([]string, 0)
	if len(pokemon.Moves) > 10 {
		tableHeader := "||||\n|---|---|---|"
		for i := 0; i < (len(pokemon.Moves)+2)/3; i++ {
			start := i * 3
			end := min(start+3, len(pokemon.Moves))

			row := make([]string, 3)
			for j := range 3 {
				if start+j < end {
					row[j] = pokemon.Moves[start+j].Move.Name
				} else {
					row[j] = "" // Pad with empty string
				}
			}
			movesLine := fmt.Sprintf(
				"|- %s | - %s | - %s |",
				utils.CapitalizeFirstLetters(row[0]),
				utils.CapitalizeFirstLetters(row[1]),
				utils.CapitalizeFirstLetters(row[2]),
			)
			movesSlice = append(movesSlice, movesLine)
		}
		return fmt.Sprintf("## Moves\n%s\n%s", tableHeader, strings.Join(movesSlice, "\n"))
	} else {
		for _, m := range pokemon.Moves {
			moveLine := fmt.Sprintf(
				"- %s",
				utils.CapitalizeFirstLetters(m.Move.Name),
			)
			movesSlice = append(movesSlice, moveLine)
		}
		return fmt.Sprintf("## Moves\n %s", strings.Join(movesSlice, "\n"))
	}
}
