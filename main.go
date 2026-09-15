package main

import (
	"bufio"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func ValidateURL(rawURL string) (*url.URL, error) {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL structure: %w", err)
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return nil, errors.New("URL scheme must be http or https")
	}

	if u.Hostname() == "" {
		return nil, errors.New("URL missing host domain")
	}

	return u, nil
}

func handleCommand(store *URLStore, cmd string, args []string) error {
	switch cmd {
	case "shorten":
		if len(args) == 0 {
			return fmt.Errorf("shorten requires a URL")
		}

		shortUrlCode := Shorten(store, args[0])
		fmt.Printf("Short Code: %s\n", shortUrlCode)
	case "expand":
		if len(args) == 0 {
			return fmt.Errorf("expand requires a short code")
		}

		longUrl, err := Expand(store, args[0])
		if err != nil {
			return err
		}
		fmt.Println(longUrl)
	case "list":
		url := store.All()
		for key, value := range url {
			fmt.Printf("%s : %s\n", key, value)
		}
	case "help":
		fmt.Print(`Commands:
		shorten <URL>     - Create a short code for a URL
		expand <code>     - Get the original URL
		list              - Show all mappings
		help              - Show this message
		exit              - Exit the program
	  `)
	default:
		return fmt.Errorf("unknown command: %s", cmd)
	}
	return nil
}

func main() {
	store := NewURLStore()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Welcome to URL Shortener!\nType 'help' for commands.\n")

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "failed to read input:", err)
				os.Exit(1)
			}
			break
		}

		input := scanner.Text()
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		command := parts[0]
		args := parts[1:]

		if command == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		if command == "shorten" && len(args) == 0 {
			fmt.Println("Error: shorten requires a URL")
			continue
		}
		if command == "expand" && len(args) == 0 {
			fmt.Println("Error: expand requires a short code")
			continue
		}

		err := handleCommand(store, command, args)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
