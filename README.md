# 🎮 Gotta Catch'em All! Go-Pokédex

> A terminal-based Pokédex built in Go — started as a guided [Boot.dev](https://boot.dev) project, evolving into a full CLI Pokémon game.

---

## About

This project started as a guided Boot.dev exercise to reinforce core Go and HTTP concepts — things like making API requests, handling JSON responses, and implementing an in-memory cache. But the plan doesn't stop there. As more Boot.dev courses get completed, this will grow into a proper CLI Pokémon game with exploration, battles, gyms, and progression.

Under the hood it uses the [PokéAPI](https://pokeapi.co/) for all Pokémon and location data.

---

## Features

- 🗺️ **Map Navigation** — Browse location areas 20 at a time, forward and backward
- 🔍 **Explore** — Discover which Pokémon inhabit a given location area
- 🎣 **Catch** — Attempt to catch Pokémon (catch rate based on base experience)
- 📋 **Inspect** — View detailed stats, types, and height/weight for caught Pokémon
- 📓 **Pokédex** — List every Pokémon you've caught so far
- ⚡ **Caching** — API responses are cached in-memory with a TTL to reduce redundant requests
- 🖥️ **REPL Interface** — Interactive command loop with a `help` command listing all available commands

---

## Commands

| Command | Description |
|---|---|
| `help` | Display available commands |
| `map` | List the next 20 location areas |
| `mapb` | Go back to the previous 20 location areas |
| `explore <area>` | List Pokémon found in a location area |
| `catch <pokemon>` | Attempt to catch a Pokémon |
| `inspect <pokemon>` | View stats for a caught Pokémon |
| `pokedex` | List all Pokémon in your Pokédex |
| `exit` | Quit the program |

---

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.21+

### Installation

```bash
git clone https://github.com/AggroSec/Go-Pokedex.git
cd Go-Pokedex
go build -o pokedex
./pokedex
```

---

## Roadmap

This project is actively being extended. Here's what's planned:

### 🔜 Near Future
- [ ] **Save progress to disk** — Persist your Pokédex between sessions so your catches aren't lost on exit
- [ ] **Command history** — Use the up arrow key to cycle through previously entered commands

### 🔭 Further Down the Road
- [ ] **Improved exploration** — Richer area discovery and movement between connected locations
- [ ] **Battle simulations** — Turn-based battles against wild Pokémon
- [ ] **Gyms** — Gym leaders to challenge with unique battle rules
- [ ] **Level ups & XP** — Pokémon gain experience and grow stronger over time
- [ ] **And more...** — As Boot.dev courses are completed, new mechanics will keep getting added

---

## Tech

- **Language:** Go
- **API:** [PokéAPI](https://pokeapi.co/)
- **Architecture:** REPL loop with command dispatch, internal HTTP client, TTL cache

---

## Acknowledgements

Built following the [Boot.dev](https://boot.dev) guided project curriculum. Boot.dev is a fantastic platform for learning backend development through hands-on projects — highly recommended.

---

*Gotta catch 'em all.*
