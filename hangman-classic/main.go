package main

import (
	"fmt"
	"strings"
	"HangmanWeb/hangman-classic/functions"
)

func main() {
	// Initialisation du jeu
	functions.Addwords()

	// Choisir un mot au hasard
	wordToGuess := functions.GetRandomWord()
	CurrentWordState := functions.InitializeWord(wordToGuess)
	CurrentWordState[0] = string(wordToGuess[0]) // Révéler la première lettre            

	guessedWords := []string{}
	attempts := 10 // Nombre initial de tentatives

	for attempts > 0 {
		// Afficher l'état actuel du jeu
		functions.DisplayState(CurrentWordState, attempts)

		// Demander à l'utilisateur de deviner
		fmt.Print("Devine la lettre ou le mot: ")
		guess := functions.InputUser()

		if guess == "" {
			fmt.Println("Veuillez entrer une lettre ou un mot.")
			continue
		}

		// Convertir la lettre en minuscule pour gérer les différences de casse
		guess = functions.ConvertToLower(guess)

		// Vérification si le guess est valide
		if !functions.IsValidInput(guess) {
			fmt.Println("Veuillez entrer uniquement des lettres.")
			continue
		}

		// Si l'utilisateur devine une lettre
		if len(guess) == 1 {
			functions.HandleLetterGuess(guess, &guessedWords, wordToGuess, CurrentWordState, &attempts)
		} else { // Si l'utilisateur devine un mot entier
			functions.HandleWordGuess(guess, &guessedWords, wordToGuess, &attempts, CurrentWordState)
		}

		// Vérifier si le mot a été deviné correctement
		if strings.Join(CurrentWordState, "") == wordToGuess {
			fmt.Println("Félicitations, vous avez deviné le mot:", wordToGuess)
			return
		}

		// Si l'utilisateur a épuisé toutes ses tentatives
		if attempts <= 0 {
			fmt.Println("GAME OVER. Le mot était:", wordToGuess)
			return
		}
	}
}
