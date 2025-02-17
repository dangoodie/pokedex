package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var cfg Config

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		args := cleanInput(scanner.Text())
		if len(args) == 0 {
			continue
		}

		command, ok := getCommands()[args[0]]
		if ok {
			err := command.callback(&cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
