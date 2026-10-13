package main

import "fmt"

func commandHelp() error {
	helpStr := `Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex

`

	fmt.Print(helpStr)
	return nil
}
