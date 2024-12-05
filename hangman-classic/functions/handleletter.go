package functions

import (
    "fmt"
    "strings"
)

// DisplayGuesses displays the current word state, remaining attempts, and guessed letters and words
func DisplayGuesses(currentWordState []string, attempts int, GuessedLetters []string, GuessedWords []string) {
    // Display the current word state (e.g., a p p _ _)
    fmt.Println("État actuel du mot :", strings.Join(currentWordState, " "))
    // Display the remaining attempts
    fmt.Println("Tentatives restantes :", attempts)

    // Display guessed letters, if any
    if len(GuessedLetters) > 0 {
        fmt.Println("Lettres déjà devinées :", strings.Join(GuessedLetters, ", "))
    } else {
        fmt.Println("Aucune lettre devinée encore.")
    }

    // Display guessed words, if any
    if len(GuessedWords) > 0 {
        fmt.Println("Mots déjà devinés :", strings.Join(GuessedWords, ", "))
    } else {
        fmt.Println("Aucun mot deviné encore.")
    }

    // Display the hangman position based on the remaining attempts
    DisplayHangman(attempts)
}

// HandleWordGuess handles the guess of a word
func HandleWordGuess(guess string, GuessedWords *[]string, wordToGuess string, attempts *int, currentWordState []string) {
    if Contains(*GuessedWords, guess) {
        fmt.Println("Vous avez déjà deviné ce mot.")
        return
    }
    *GuessedWords = append(*GuessedWords, guess)
    if guess == wordToGuess {
        fmt.Println("Félicitations! Vous avez deviné le mot!")
        return
    } else {
        *attempts -= CalculateAttemptLoss(guess) // Subtract attempts based on the word length
        fmt.Println("Mauvais mot!")
        DisplayHangman(*attempts) // Display hangman position
    }
}

// CalculateAttemptLoss calculates the number of attempts lost based on the guessed word
func CalculateAttemptLoss(guess string) int {
    // Subtract attempts equal to the length of the guessed word
    return len(guess)
}

// HandleLetterGuess handles the guess of a single letter
func HandleLetterGuess(guess string, GuessedLetters *[]string, wordToGuess string, currentWordState []string, attempts *int) {
    if Contains(*GuessedLetters, guess) {
        fmt.Println("Vous avez déjà deviné cette lettre.")
        return
    }
    *GuessedLetters = append(*GuessedLetters, guess)
    if UpdateWord(wordToGuess, currentWordState, guess) {
        fmt.Println("Bonne lettre!")
    } else {
        *attempts-- // Subtract one attempt for a wrong letter guess
        fmt.Println("Mauvaise lettre!")
        DisplayHangman(*attempts) // Display hangman position
    }
}

// LoadHangmanPositionsAndDisplay displays the hangman drawing based on the remaining attempts
