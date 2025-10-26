# CLI Pokedex in Go

A fun command-line tool to fetch and display Pokémon information, including sprites, cry sounds, and descriptions. All in your terminal!

---

## Features
- Fetch Pokémon by name or ID.
- Display sprites in the terminal using [`chafa`](https://github.com/hpjansson/chafa).
- Play Pokémon cry sounds.
- Show textual descriptions.

---

## Requirements
- Go 1.20+
- [`chafa`](https://github.com/hpjansson/chafa) (for terminal image rendering)
- A default audio player for cry sounds (on Linux I chose to use `sox`)

### Install `chafa`
#### Linux (Debian/Ubuntu)
```bash
sudo apt-get install chafa
```
#### macOS (Homebrew)
```bash
brew install chafa
```
#### Windows
Download from [chafa releases](https://github.com/hpjansson/chafa/releases) or use WSL.

### Install `sox` to play cry sound on Linux
```bash
sudo apt-get install sox
```
---

## Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/blackpantech/pokedex-cli.git
    cd pokedex-cli
    ```
2. Build the app:
    ```bash
    go build -o pokedex-cli
    ```
3. Move the binary to your PATH (optional):
    ```bash
    mv pokedex-cli /your/local/bin/path/
    ```

---

## Usage
Run the app with a Pokémon name or ID:
```bash
pokedex-cli pikachu
# or
pokedex-cli 25
```
Example output:
![screenshot 1](/screenshots/screenshot-1.png)
![screenshot 2](/screenshots/screenshot-2.png)

---

## API
This app uses the [PokéAPI](https://pokeapi.co/), a free RESTful API for Pokémon data.

---

## Why This Project?
I built this app to practice **Go** by combining:
- **API usage** (fetching data from PokéAPI)
- **File handling** (downloading and managing sprites/sounds)
- **Command-line interface design** (using Cobra for a user-friendly CLI and Glamour for a nicer render in the terminal)

It was a fun way to explore Go’s standard library and external tools like `chafa` while creating something useful for Pokémon fans!
