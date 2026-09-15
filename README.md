# URL Shortener CLI

A command-line tool that creates short codes for long URLs and allows you to look them up later.

## Features

- Generate unique 6-character short codes for URLs
- Look up original URLs by their short codes
- Thread-safe concurrent access with `sync.Mutex`
- Interactive CLI interface

## Installation

```bash
git clone https://github.com/Eyob49/url-shortener.git
cd url-shortener
go build -o url-shortener
```

## Usage

```bash
$ go run main.go
Welcome to URL Shortener!
Type 'help' for commands.

> shorten https://www.wikipedia.org/wiki/Go_(programming_language)
Short Code: gURK90

> expand gURK90
https://www.wikipedia.org/wiki/Go_(programming_language)

> list
gURK90 : https://www.wikipedia.org/wiki/Go_(programming_language)

> exit
Goodbye!
```

## Commands

- `shorten <URL>` - Create a short code for a URL
- `expand <code>` - Get the original URL
- `delete <code>` - Delete a short code mapping
- `list` - Show all mappings
- `help` - Show available commands
- `exit` - Exit the program