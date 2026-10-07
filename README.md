# Pokedex CLI

A small command-line Pokedex I made while learning about HTTP in Go. It uses the [PokéAPI](https://pokeapi.co/) to look up locations and Pokémon.

## Run

You need Go 1.27.1 or newer and an internet connection to make requests to the API.

```sh
go run .
```

The program starts an interactive prompt:

```text
Pokedex >
```

## Commands

| Command | Description |
| --- | --- |
| `help` | Show the available commands. |
| `map` | Show the next page of locations. |
| `mapb` | Show the previous page of locations. |
| `explore <area-name>` | List Pokémon found in a location area. |
| `catch <pokemon-name>` | Try to catch a Pokémon. |
| `inspect <pokemon-name>` | Show details for a Pokémon caught in this session. |
| `pokedex` | List Pokémon caught in this session. |
| `exit` | Quit the program. |

Caught Pokémon are kept in memory and are lost when the program exits. API responses are cached while the program is running.
