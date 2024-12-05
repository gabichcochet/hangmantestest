package Hangman

import "strings"

// Contains vérifie si un slice contient un élément spécifique.
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// Contains vérifie si un slice contient un élément spécifique.

// FindLetter vérifie si une lettre est présente dans un mot donné.
func FindLetter(guessedLetters []string, letter string) bool {
    for _, guessed := range guessedLetters {
        if guessed == letter {
            return true // Letter has already been guessed
        }
    }
    return false // Letter has not been guessed yet
}

// FindLetter vérifie si une lettre est présente dans un mot donné.

// IsValidInput checks if the input contains only valid letters (a-z, A-Z)
func IsValidInput(guess string) bool {
	for _, char := range guess {
		if !('a' <= char && char <= 'z') && !('A' <= char && char <= 'Z') {
			return false
		}
	}
	return true
}

// ConvertToLower converts the input guess to lowercase
func ConvertToLower(guess string) string {
	return strings.ToLower(guess)
}
