# serve

A lightweight utility for creating a local server from a directory.

## Building from Source

This Readme assumes you're using [task](https://taskfile.dev/docs/installation).

### Building

Simply run `task build` to build. The binaries will appear in the `out/` directory.

You can also build for your specific platform with:

- `task build-darwin` for macOS
- `task build-linux` for Linux
- `task build-windows` for Windows

## Dependencies

- [Lip Gloss](https://github.com/charmbracelet/lipgloss)
- [x/term](https://pkg.go.dev/golang.org/x/term)
- [keyboard](https://github.com/eiannone/keyboard)
