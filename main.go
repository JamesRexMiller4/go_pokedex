package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	internal "github.com/jamesrexmiller4/go_pokedex/internal/apis"
)

func main() {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	var config *Config
	for {
		for scanner.Scan() {
			cmd := scanner.Text()
			fmt.Printf("Pokedex > %s\n", cmd)

			switch cmd {
			case "help":
				helpCmd(config)
			case "exit":
				exitCmd(config)
			case "map":
				mapCmd(config)
			default:
				fmt.Println("command not recognized")
			}
		}
	}
}

func commandsMap() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCmd,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    exitCmd,
		},
		"map": {
			name:        "map",
			description: "Get next page of results",
			callback:    mapCmd,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous page of results",
			callback:    mapbCmd,
		},
	}
}

func helpCmd(config *Config) error {
	fmt.Println("\nWelcome to the Pokedex")
	fmt.Println("Usage:")

	for key := range commandsMap() {
		obj := commandsMap()[key]
		fmt.Printf("%s: %s \n", obj.name, obj.description)
	}
	fmt.Print("\nPokedex >")

	return nil
}

func exitCmd(config *Config) error {
	os.Exit(0)
	return nil
}

func mapCmd(config *Config) error {
	// Maps pagination results forward
	next := &config.Next
	fmt.Printf("Next: %d", next)
	url := "https://pokeapi.co/api/v2/location?offset=0&limit=20"

	res, err := internal.GetLocation(url)
	if err != nil {
		return fmt.Errorf("failed to get: %w", err)
	}

	if err = json.Unmarshal(res, &config); err != nil {
		return fmt.Errorf("unable to unmarshal response: %w", err)
	}

	logResults(config)

	fmt.Printf("Next after fetch: %s", config.Next)

	return nil
}

func mapbCmd(config *Config) error {
	// Maps pagination results backward
	return nil
}

func logResults(config *Config) {
	for _, result := range config.Results {
		fmt.Printf("%s\n", result.Name)
	}
}
