# Pokedex CLI

A command-line Pokédex written in Go. Explore the Pokémon world region by region, catch Pokémon, inspect their stats, and build up your own personal collection — all from your terminal, powered by [PokeAPI](https://pokeapi.co/).

## Features

- **Explore the map** — page forward and backward through location areas
- **Discover Pokémon** — see which Pokémon appear in a given area
- **Catch Pokémon** — throw a Pokéball with a catch chance based on each Pokémon's base experience
- **Inspect your catches** — view height, weight, stats, and types for any Pokémon you've caught
- **Track your Pokédex** — list every Pokémon you've caught so far
- **Built-in caching** — API responses are cached in memory with automatic expiry, so repeated lookups don't hammer PokeAPI

## Requirements

- [Go](https://go.dev/dl/) 1.26 or later

## Installation

Clone the repository and build the binary:

```bash
git clone https://github.com/SuperJake03/pokedex-cli.git
cd pokedex-cli
go build -o pokedex
```

Then run it:

```bash
./pokedex
```

Alternatively, run it directly without building a binary:

```bash
go run .
```

## Usage

Once running, you'll get an interactive REPL prompt:

```
Pokedex >
```

### Commands

| Command | Description |
|---|---|
| `help` | Displays a help message listing all commands |
| `map` | Get the next page of location areas |
| `mapb` | Get the previous page of location areas |
| `explore <location_name>` | List all Pokémon found in a given location area |
| `catch <pokemon_name>` | Attempt to catch a Pokémon |
| `inspect <pokemon_name>` | View stats, types, height, and weight for a caught Pokémon |
| `pokedex` | List every Pokémon you've caught |
| `exit` | Exit the Pokedex |

### Example session

```
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokemon:
 - tentacool
 - magikarp
 - gyarados

Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
magikarp was caught!
You may now inspect it with the inspect command.

Pokedex > inspect magikarp
Name: magikarp
Height: 9
Weight: 100
Stats:
  -hp: 20
  -attack: 10
  -defense: 55
  -special-attack: 15
  -special-defense: 20
  -speed: 80
Types:
  - water

Pokedex > pokedex
Your Pokedex:
 - magikarp
```

## Project structure

```
.
├── main.go                  # Entry point; wires up the API client and config
├── repl.go                  # REPL loop and input parsing
├── commands.go               # Command definitions and their callbacks
├── repl_test.go              # Tests for input parsing
├── go.mod
└── internal/
    ├── pokeapi/              # PokeAPI client and response types
    │   ├── client.go         # HTTP client with caching wrapper
    │   ├── pokeapi.go        # Base API URL
    │   ├── pokemon.go        # Pokémon lookup + response types
    │   ├── pokemon_list.go   # Location area Pokémon encounters
    │   └── locations_list.go # Paginated location area listing
    └── pokecache/            # In-memory cache with time-based expiry
        ├── cache.go
        └── cache_test.go
```

## How it works

- **API client** (`internal/pokeapi`): wraps `net/http` to talk to PokeAPI, decoding JSON responses into typed Go structs. Every request is routed through the cache first.
- **Cache** (`internal/pokecache`): a simple thread-safe, in-memory key/value store keyed by request URL. A background reaper loop periodically clears out entries older than the configured interval, so cached data doesn't grow unbounded or get stale forever.
- **REPL** (`repl.go`): reads a line of input, splits it into a command and arguments, and dispatches to the matching command's callback.
- **Catch mechanic** (`commands.go`): catch probability is calculated as `150 / (150 + baseExperience)`, so lower base-experience Pokémon are easier to catch and higher base-experience Pokémon are tougher.

## Testing

Run the test suite with:

```bash
go test ./...
```
