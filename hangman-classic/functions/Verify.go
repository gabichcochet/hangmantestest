package functions

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
func FindLetter(word string, letter string) bool {
	for _, char := range word {
		if string(char) == letter {
			return true
		}
	}
	return false
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
