package main

import (
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
