# TLF - tiny little finder 

TLF is a CLI clipboard history manager that allows you to save and retrieve key, [value, link] pairs. It helps you manage your clipboard history efficiently.

## Quick Start

### Install TLF using Homebrew on macOS
```bash
$ brew tap thegeorgejoseph/homebrew-tap
$ brew install tlf
```

## Usage

### Set Command
Use the set command to save a key-value pair with an optional link to the clipboard history.

```bash
$ tlf set -k myKey -v myValue -l myLink
```

### Get Command
Use the get command to retrieve the value or link of a key in the clipboard history.

```bash
# Get the value of a key
$ tlf get myKey -v

# Get the link of a key
$ tlf get myKey -l

# Automatically copy the response to the clipboard
$ tlf get myKey -v
```

### Additional Commands
- help: Display help information about any command.

## Flags
- -h, --help: Display help for the main command.
- -k, --key: Key for the set command.
- -v, --value: Value for the set command or to specify response type in get command.
- -l, --link: Link for the set command or to specify response type in get command.

## Supported Operating Systems
Use the table below to see the supported operating systems:

| OS       | Architectures           | Supported |
|----------|-------------------------|-----------|
| macOS    | Intel, Apple Silicon    |  Yes      |
| Windows  | 32-bit, 64-bit, ARM64   |  Yes      |
| Linux    | 32-bit, 64-bit, ARM64   |  Yes      |

## Contributing

Feel free to contribute to TLF by opening a PR.

## License

This project is licensed under the BSD License.





