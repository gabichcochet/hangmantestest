package Hangman

import (
	"fmt"
)

// HangmanPositions contient les différentes étapes du pendu sous forme de tableau de chaînes
var HangmanPositions = []string{
	"",
	"=========",
	"=========\n|\n|\n|\n|\n|",
	"=========\n|\n|\n|\n|\n|\n===========",
	"=========\n|/\n|\n|\n|\n|\n===========",
	"=========\n|/  |\n|\n|\n|\n|\n===========",
	"=========\n|/  |\n|   O\n|\n|\n|\n===========",
	"=========\n|/  |\n|   O\n|  /|\n|\n|\n===========",
	"=========\n|/  |\n|   O\n|   |\n|\n|\n===========",
	"=========\n|/  |\n|   O\n|  /|\\\n|  /\n|\n===========",
	"=========\n|/  |\n|   O\n|  /|\\\n|  / \\\n|\n===========",
}

// UpdateHangmanPosition imprime la position actuelle du pendu en fonction du nombre de tentatives restantes.
func UpdateHangmanPosition(attempts int) {
	// On calcule l'index du sous-ensemble des lignes à afficher
	positionIndex := 10 - attempts // Si 10 tentatives sont possibles, indexera de 0 à 9

	// Limiter l'index pour ne pas sortir des bornes du tableau
	if positionIndex < 0 {
		positionIndex = 0
	}

	// Afficher l'état correspondant aux tentatives restantes
	fmt.Println(HangmanPositions[positionIndex])
}
