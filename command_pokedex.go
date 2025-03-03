package main

import "fmt"

func commandPokedex(cfg *config, _ ...string) error {
	fmt.Println("Your Pokedex:")
	for _, p := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
