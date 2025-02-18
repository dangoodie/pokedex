# Pokedex CLI

A simple command-line Pokédex powered by the [PokéAPI](https://pokeapi.co/). This tool allows you to explore, catch, and inspect Pokémon using a lightweight terminal interface.

## Features
- **Explore Locations**: Search for Pokémon in different locations.
- **Catch Pokémon**: Attempt to catch Pokémon based on their base experience.
- **Inspect Pokémon**: View details of the Pokémon you have caught.
- **View Pokedex**: List all Pokémon you have caught.
- **Navigation**: Move forward or backward on the map.

## Commands
| Command  | Description |
|----------|------------|
| `help`   | Display a help message. |
| `exit`   | Exit the Pokédex. |
| `map`    | Move forward on the map 20 locations. |
| `mapb`   | Move backward on the map 20 locations. |
| `explore` | Search a location for Pokémon. Usage: `explore <location-name>` |
| `catch`  | Try to catch a Pokémon. Usage: `catch <pokemon-name>` |
| `inspect` | Inspect a Pokémon that you have caught. Usage: `inspect <pokemon-name>` |
| `pokedex` | List all Pokémon in your Pokédex. |

## How It Works
- **Catching Pokémon**: The probability of catching a Pokémon is based on its base experience (XP). Stronger Pokémon are harder to catch.
- **PokéAPI Integration**: All Pokémon data is retrieved from the [PokéAPI](https://pokeapi.co/), ensuring accurate and up-to-date information.

## Installation & Usage
1. Clone this repository:
   ```sh
   git clone https://github.com/dangoodie/pokedex.git
   cd pokedex-cli
   ```
2. Run the CLI:
   ```sh
   go run .
   ```

---
Built with ❤️ using Go and the PokéAPI.

