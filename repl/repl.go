package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/trantuvan/pokedex-cli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(filter string) error
}

type config struct {
	Prev       *string //* allow nil string
	Next       string
	httpClient pokeapi.Client
}

func getCliCommands() map[string]cliCommand {
	urls := &config{httpClient: pokeapi.NewClient(5*time.Second, 5*time.Minute)}

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback: func(string) error {
				return commandHelp()
			},
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback: func(filter string) error {
				return commandExit()
			},
		},
		"map": {
			name:        "map",
			description: "displays the names of the next 20 location areas in the Pokemon world",
			callback: func(string) error {
				return commadMap(urls)
			},
		},
		"mapb": {
			name:        "mapb",
			description: "displays the names of the previous 20 location areas in the Pokemon world",
			callback: func(string) error {
				return CommandMapb(urls)
			},
		},
		"explore": {
			name:        "explore",
			description: "displays the names of pokemons in this area",
			callback: func(filter string) error {
				return CommandExplore(urls, filter)
			},
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func Repl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCliCommands()

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		if command, exist := commands[words[0]]; exist {
			filter := ""
			if len(words) == 2 {
				filter = words[1]
			}
			err := command.callback(filter)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown commad")
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error in standard in: %v", err)
	}
}
