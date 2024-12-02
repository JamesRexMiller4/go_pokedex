package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Pokedex >")
		for scanner.Scan() {
			cmd := scanner.Text()

			switch cmd {
			case "help":
				help()
			case "exit":
				exit()
			default:
				fmt.Println("command not recognized")
			}
		}
	}
}

func commandsMap() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    help,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    exit,
		},
	}
}

func help() error {
	fmt.Println("\nWelcome to the Pokedex")
	fmt.Println("Usage:")

	for key := range commandsMap() {
		obj := commandsMap()[key]
		fmt.Printf("%s: %s \n", obj.name, obj.description)
	}

	fmt.Println("\nPokedex >")
	return nil
}

func exit() error {
	os.Exit(0)
	return nil
}
