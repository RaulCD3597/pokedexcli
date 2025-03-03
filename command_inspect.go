package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	if _, ok := cfg.caughtPokemon[name]; !ok {
		return errors.New("you have not caught that pokemon")
	}
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Println(pokemon)
	return nil
}
