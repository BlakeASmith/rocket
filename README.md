# Rocket

Rocket is a CLI tool for managing your project folders.

## Installation

Run `./install.sh` to build, install the binary, and set up shell integration automatically. Requires Go and optionally sudo for system installation.

## Requirements

- Go (for building)
- [fzf](https://github.com/junegunn/fzf) (for fuzzy search in goto and rm commands)

## Configuration

Rocket uses a YAML config file at `~/.config/rocket/config.yml`. The default configuration is:

```yaml
rocket_root: ~/rocket
```

You can modify `rocket_root` to point to your desired projects directory.

## Commands

### `rocket new <name>`

Creates a new project directory under the `rocket_root`.

Example:
```
rocket new myproject
```

### `rocket goto`

Opens a fuzzy search interface using `fzf` to select a subdirectory under `rocket_root`. Once selected, changes the shell's current directory to the chosen path.

Requirements: `fzf` must be installed.

Example:
```
rocket goto
```

This will list project directories and allow you to fuzzy search and select one to navigate to.

### `rocket init`

Prints the shell integration code for installation.

### `rocket ls [query]`

Lists project directories under `rocket_root`. If a query is provided, lists only directories that fuzzy match the query.

Example:
```
rocket ls
rocket ls myproject
```

### `rocket mv <source_query> <dest>`

Moves or renames a project directory using fuzzy search for the source. The destination must be relative to `rocket_root`.

Example:
```
rocket mv oldname newname
```

### `rocket rm [query]`

Removes a project directory using fuzzy search. If no query, uses fzf for selection.

Example:
```
rocket rm myproject
```

## License

See LICENSE file.