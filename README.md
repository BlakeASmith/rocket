# Rocket

Rocket is a CLI tool for managing your project folders.

## Installation

1. Build the binary: `go build -o rocket-bin`
2. Run `rocket init` to get the shell integration code.
3. Add the output to your shell profile (e.g., `~/.bashrc` or `~/.zshrc`).

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

## License

See LICENSE file.