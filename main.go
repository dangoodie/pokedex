package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		args := cleanInput(scanner.Text())
		if len(args) != 0 {
			fmt.Printf("Your command was: %s\n", args[0])
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
