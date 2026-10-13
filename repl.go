package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	var textOut []string
	var textSplit []string = strings.Fields(text)
	for i := range textSplit {
		textOut = append(textOut, strings.ToLower(textSplit[i]))
	}
	return textOut
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			// Scan returns true
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "error reading input: %v", err)
				os.Exit(0)
			}
			return
		}
		words := cleanInput(scanner.Text())
		commandName := words[0]
		if len(words) == 0 {
			continue
		}
		command, exists := commands[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
