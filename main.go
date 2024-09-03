package main

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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCliCommands()

	for {
		fmt.Print("Pokedex >")

		if !scanner.Scan() {
			break
		}

		key := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if key == "help" {
			commands[key].callback()
			continue
		} else if key == "exit" {
			commands[key].callback()
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error in standard in: %v", err)
	}
}

func commandHelp(commands map[string]cliCommand) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n")

	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func() error { return commandHelp(getCliCommands()) },
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    func() error { return nil },
		},
	}
}
