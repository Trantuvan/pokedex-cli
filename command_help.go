package main

import "fmt"

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n")

	for _, v := range getCommands() {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}
