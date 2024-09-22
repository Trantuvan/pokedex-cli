package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	Next string
	Prev *string //* allow nil string
}

func getCliCommands() map[string]cliCommand {
	urls := &config{}

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "displays the names of the next 20 location areas in the Pokemon world",
			callback: func() error {
				return commadMap(urls)
			},
		},
		"mapb": {
			name:        "mapb",
			description: "displays the names of the previous 20 location areas in the Pokemon world",
			callback: func() error {
				return CommandMapb(urls)
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
		fmt.Print("Pokedex >")

		if !scanner.Scan() {
			break
		}

		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		if command, exist := commands[words[0]]; exist {
			err := command.callback()
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
