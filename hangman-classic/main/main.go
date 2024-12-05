package main

import (
	"fmt"
	"strings"
	Hangman "Hangman/functions" // Importing the functions package
)

func main() {
	// Initialisation du jeu
	Hangman.Addwords("./txt/words.txt") // Adding words from the file

	// Choisir un mot au hasard
	wordToGuess := Hangman.GetRandomWord() // Get a random word to guess
	currentWordState := Hangman.InitializeWord(wordToGuess) // Initialize the word state
	currentWordState[0] = string(wordToGuess[0]) // Reveal the first letter            

	guessedLetters := []string{} // Store guessed letters
	guessedWords := []string{} // Store guessed words
	attempts := 10 // Initial number of attempts

	// Game loop, continue until attempts are used up or the word is guessed
	for attempts > 0 {
		// Display the current state of the game
		Hangman.DisplayGuesses(currentWordState, attempts, guessedLetters, guessedWords)

		// Ask the user to guess a letter or a word
		fmt.Print("Devine la lettre ou le mot: ")
		guess := Hangman.InputUser() // Get the user's input

		// Check for invalid input
		if guess == "" {
			fmt.Println("Veuillez entrer une lettre ou un mot.")
			continue
		}

		// Convert the guess to lowercase to handle case sensitivity
		guess = Hangman.ConvertToLower(guess)

		// Validate the input (check if it's a letter or word)
		if !Hangman.IsValidInput(guess) {
			fmt.Println("Veuillez entrer uniquement des lettres.")
			continue
		}

		// Process the guess
		if len(guess) == 1 {
			// If the user guessed a letter
			Hangman.HandleLetterGuess(guess, &guessedLetters, wordToGuess, currentWordState, &attempts)
		} else {
			// If the user guessed a word
			Hangman.HandleWordGuess(guess, &guessedWords, wordToGuess, &attempts, currentWordState)
		}

		// Check if the word has been guessed correctly
		if strings.Join(currentWordState, "") == wordToGuess {
			fmt.Println("Félicitations, vous avez deviné le mot:", wordToGuess)
			return
		}

		// If the user has run out of attempts
		if attempts <= 0 {
			fmt.Println("GAME OVER. Le mot était:", wordToGuess)
			return
		}
	}
}