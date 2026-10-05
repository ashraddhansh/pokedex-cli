package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowercase := strings.ToLower(text)
	trimmed := strings.Fields(lowercase)
	return trimmed
}


