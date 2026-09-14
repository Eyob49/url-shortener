package main

import (
	"fmt"
	"math/rand/v2"
)

func GenerateShortCode(length int) string {
	const alphChar = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, length)

	for i := 0; i < length; i++ {
		randomIndex := rand.IntN(len(alphChar))
		result[i] = alphChar[randomIndex]
	}

	return string(result)
}

func Shorten(store *URLStore, longURL string) string {
	for {
		randomToken := GenerateShortCode(6)
	_, exists := store.Get(randomToken)
	if !exists {
		store.Add(randomToken, longURL)
		return randomToken
	}

	}
	

	
}

func Expand(store *URLStore, shortCode string) string {
	longURL, exists := store.Get(shortCode)
	if exists {
		return fmt.Printf("%s : %s", shortCode, longURL)
	}
	fmt.Errorf("No matching record found")
	return nil
}
