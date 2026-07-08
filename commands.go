package main

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
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
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Get all Pokémon located in <area-name>",
			callback:    commandExplore,
		},
	}
}

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *config, args []string) error {
	locationAreas, err := cfg.pokeapiClient.ListLocationsAreas(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}
	cfg.nextLocationsURL = locationAreas.Next
	cfg.prevLocationsURL = locationAreas.Previous
	return nil
}

func commandMapb(cfg *config, args []string) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locationAreas, err := cfg.pokeapiClient.ListLocationsAreas(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	cfg.nextLocationsURL = locationAreas.Next
	cfg.prevLocationsURL = locationAreas.Previous

	return nil
}

func commandExplore(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("location name required!! Usage: explore <location-name>")
	}

	areaName := args[0]

	area, err := cfg.pokeapiClient.ListAreaPokemon(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")

	for _, pokeEncounter := range area.PokemonEncounters {
		pokemonName := pokeEncounter.Pokemon.Name
		fmt.Printf(" - %s\n", pokemonName)
	}

	return nil
}
