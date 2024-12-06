package Hangman

import (
    "fmt"
    "strings"
)

func DisplayGuesses(currentWordState []string, attempts int, guessedLetters []string, guessedWords []string) {
	// Display the current word state (e.g., a p p _ _)
	fmt.Println("État actuel du mot :", strings.Join(currentWordState, " "))
	// Display the remaining attempts
	fmt.Println("Tentatives restantes :", attempts)

	// Display guessed letters, if any
	if len(guessedLetters) > 0 {
		fmt.Println("Lettres déjà devinées :", strings.Join(guessedLetters, ", "))
	} else {
		fmt.Println("Aucune lettre devinée encore.")
	}

	// Display guessed words, if any
	if len(guessedWords) > 0 {
		fmt.Println("Mots déjà devinés :", strings.Join(guessedWords, ", "))
	} else {
		fmt.Println("Aucun mot deviné encore.")
	}
}


func HandleWordGuess(guess string, guessedWords *[]string, wordToGuess string, attempts *int, currentWordState []string) {
	// Check if the word is too long for the remaining attempts
	if len(guess) > *attempts {
		*attempts = 0 // Set attempts to 0, triggering GAME OVER
	} else {
		// Regular handling of word guesses
		if Contains(*guessedWords, guess) {
			fmt.Println("Vous avez déjà deviné ce mot.")
			return
		}
		*guessedWords = append(*guessedWords, guess)
		if guess == wordToGuess {
			fmt.Println("Félicitations! Vous avez deviné le mot!")
			return
		} else {
			*attempts -= CalculateAttemptLoss(guess) // Subtract attempts based on the word length
			fmt.Println("Mauvais mot!")
			// Use the Hangman package to update the hangman position
			UpdateHangmanPosition(*attempts) // Display the current hangman position
		}
	}
}


func CalculateAttemptLoss(guess string) int {
	// Subtract attempts equal to the length of the guessed word
	return len(guess)
}


func HandleLetterGuess(guess string, guessedLetters *[]string, wordToGuess string, currentWordState []string, attempts *int) {
	if Contains(*guessedLetters, guess) {
		fmt.Println("Vous avez déjà deviné cette lettre.")
		return
	}
	*guessedLetters = append(*guessedLetters, guess)
	if UpdateWord(wordToGuess, currentWordState, guess) {
		fmt.Println("Bonne lettre!")
	} else {
		*attempts-- // Subtract one attempt for a wrong letter guess
		fmt.Println("Mauvaise lettre!")
		LoadHangmanPositionsAndDisplay(*attempts) // Call LoadHangmanPositions
	}
}