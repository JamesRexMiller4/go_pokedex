package main

import "strings"

func cleanInput(input string) []string {
	lowered := strings.ToLower(input)
	words := strings.Fields(lowered)
	return words
}
