package Hangman

import (
	"fmt"
)

// InitializeWord crée un slice représentant l'état actuel du mot, rempli de underscores.
func InitializeWord(word string) []string {
	// Crée une slice avec la même longueur que le mot à deviner
	currentWordState := make([]string, len(word))
	// Remplir chaque élément avec un underscore "_"
	for i := range currentWordState {
		currentWordState[i] = "_"
	}
	return currentWordState
}

// InitializeWord crée un slice représentant l'état actuel du mot, rempli de underscores.

// initialisation et boucle principale
func InitMainLoop() {
	Addwords("./txt/words.txt")

	// Récupérer un mot aléatoire à deviner
	wordToGuess := GetRandomWord()

	// Initialiser l'état actuel du mot avec des underscores
	currentWordState := InitializeWord(wordToGuess)

	// Définir la première lettre du mot
	currentWordState[0] = string(wordToGuess[0])

	// Supposons que 'guess' est une chaîne de caractères entrée par l'utilisateur
	var guess string
	valid := true

	// Vérifie si les caractères entrés sont des lettres
	for _, char := range guess {
		if !('a' <= char && char <= 'z') && !('A' <= char && char <= 'Z') {
			valid = false
			fmt.Println("Caractère invalide détecté:", string(char))
			return
		}
	}

	// Si 'valid' est vrai, on peut continuer le jeu
	if valid {
		fmt.Println("Les caractères entrés sont valides.")
		// Ici, vous pouvez ajouter plus de logique pour gérer les essais du joueur, etc.
	}
}

// initialisation et boucle principale

// Fonction pour convertir la lettre en minuscule
