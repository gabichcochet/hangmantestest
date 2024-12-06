package Hangman

import (
	"fmt"
	"os"
	"bufio"
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

func LoadHangmanPositions(filename string) ([]string, error) {
	file, _ := os.Open(filename)
	var positions []string
	var position string
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		position += scanner.Text() + "\n"
		lineCount++
		if lineCount == 8 { // 8 lines per hangman stage
			positions = append(positions, position)
			position = ""
			lineCount = 0
		}
	}
	return positions, nil
}

func LoadHangmanPositionsAndDisplay(attempts int) {
	// Ensure HangmanPositions is loaded before use
	if len(HangmanPositions) == 0 {
		var err error
		HangmanPositions, err = LoadHangmanPositions("hangman.txt")
		if err != nil {
			fmt.Println("Erreur de chargement des positions du pendu :", err)
			return
		}
	}

	if len(HangmanPositions) > 10-attempts {
		hangmanPosition := HangmanPositions[10-attempts]
		fmt.Println(hangmanPosition)
	}
}