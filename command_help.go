package main

import (
	"fmt"
	"strings"
)

func commandHelp(cfg *config, args *[]string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("--- Available Commands ---")
	for _, command := range getCommands() {
		descLines := strings.Split(command.description, "\n")
		fmt.Printf("%-10s %s\n", command.name, descLines[0])

		for _, line := range descLines[1:] {
			fmt.Printf("           %s\n", line)
		}
	}
	fmt.Println()
	return nil
}
