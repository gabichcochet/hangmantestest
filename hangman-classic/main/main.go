package main

import (
	"fmt"
	"strings"
	Hangman "Hangman/functions" // Make sure the package path is correct
)

func main() {
	// Load words
	Hangman.Addwords("./txt/words.txt")

	// Get a random word to guess
	wordToGuess := Hangman.GetRandomWord()

	// Initialize the word state (with underscores)
	currentWordState := Hangman.InitializeWord(wordToGuess)

	// Track guessed words and letters
	guessedWords := []string{}
	guessedLetters := []string{} // Initially empty, so no guesses have been made yet
	attempts := 10

	// Flag to track if the first wrong guess has been made
	firstWrongGuessMade := false

	// Game loop
	for attempts > 0 {
		// Display the current game state
		Hangman.DisplayGuesses(currentWordState, attempts, guessedLetters, guessedWords)

		// Prompt the user to guess a letter or a word
		fmt.Print("Devine la lettre ou le mot: ")
		guess := Hangman.InputUser()
		if guess == "" {
			fmt.Println("Veuillez entrer une lettre ou un mot.")
			continue
		}

		// Validate if the input contains only letters
		if !Hangman.IsValidInput(guess) {
			fmt.Println("Veuillez entrer uniquement des lettres.")
			continue
		}

		// If the guess is a single letter
		if len(guess) == 1 {
			// Check if the letter has already been guessed
			if Hangman.Contains(guessedLetters, guess) {
				fmt.Println("Vous avez déjà deviné cette lettre.")
				continue
			}
			guessedLetters = append(guessedLetters, guess)

			// Update the current word state
			if Hangman.UpdateWord(wordToGuess, currentWordState, guess) {
				fmt.Println("Bonne lettre!")
			} else {
				attempts--  // Reduce attempts for an incorrect letter guess
				fmt.Println("Mauvaise lettre!")
				// Use the Hangman package to update the hangman position only if a wrong guess is made
				if !firstWrongGuessMade {
					firstWrongGuessMade = true
				}
				Hangman.UpdateHangmanPosition(attempts)  // Display the current hangman position
			}
		} else {
			// Handle word guesses
			Hangman.HandleWordGuess(guess, &guessedWords, wordToGuess, &attempts, currentWordState)
		}

		// Check if the current word state matches the word to guess
		if strings.Join(currentWordState, "") == wordToGuess {
			fmt.Println("Félicitations, vous avez deviné le mot:", wordToGuess)
			return
		}

		// If the player runs out of attempts
		if attempts <= 0 {
			fmt.Println("GAME OVER. Le mot était:", wordToGuess)
			// Display final hangman position
			Hangman.UpdateHangmanPosition(attempts)
			return
		}
	}
}
