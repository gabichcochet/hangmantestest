package Hangman

import (
    "fmt"
    "strings"
)

func DisplayGuesses(currentWordState []string, attempts int, GuessedLetters []string, GuessedWords []string) {
    fmt.Println("État actuel du mot :", strings.Join(currentWordState, " "))
    fmt.Println("Tentatives restantes :", attempts)

    if len(GuessedLetters) > 0 {
        fmt.Println("Lettres déjà devinées :", strings.Join(GuessedLetters, ", "))
    } else {
        fmt.Println("Aucune lettre devinée encore.")
    }

    if len(GuessedWords) > 0 {
        fmt.Println("Mots déjà devinés :", strings.Join(GuessedWords, ", "))
    } else {
        fmt.Println("Aucun mot deviné encore.")
    }

    DisplayHangman(attempts)
}

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
        *attempts -= CalculateAttemptLoss(guess)
        fmt.Println("Mauvais mot!")
        DisplayHangman(*attempts) 
    }
}


func CalculateAttemptLoss(guess string) int {
    return len(guess)
}

func HandleLetterGuess(guess string, GuessedLetters *[]string, wordToGuess string, currentWordState []string, attempts *int) {
    if Contains(*GuessedLetters, guess) {
        fmt.Println("Vous avez déjà deviné cette lettre.")
        return
    }
    *GuessedLetters = append(*GuessedLetters, guess)
    if UpdateWord(wordToGuess, currentWordState, guess) {
        fmt.Println("Bonne lettre!")
    } else {
        *attempts-- 
        fmt.Println("Mauvaise lettre!")
        DisplayHangman(*attempts) 
    }
}
