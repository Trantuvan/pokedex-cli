package repl

import "fmt"

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n")

	for _, v := range getCliCommands() {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}
