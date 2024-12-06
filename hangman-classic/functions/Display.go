package Hangman

import (
	"fmt"
	"strings"
)

// DisplayState affiche l'état actuel du mot et le nombre de tentatives restantes.
func DisplayState(currentWordState []string, attempts int, guessedLetters []string, guessedWords []string) {
	fmt.Println("État actuel du mot :", strings.Join(currentWordState, " "))
	fmt.Println("Tentatives restantes :", attempts)
	fmt.Println("Lettres déjà devinées :", strings.Join(guessedLetters, ", "))
	fmt.Println("Mots déjà devinés :", strings.Join(guessedWords, ", "))
}


// DisplayState imprime l'état actuel du mot et le nombre de tentatives restantes.

// displayHangman affiche la position du pendu selon le nombre de tentatives restantes
func DisplayHangman(attempts int) {
	positionIndex := 10 - attempts // Calcule l'index de la position à afficher

	// Si l'index est valide, afficher la position
	if positionIndex >= 0 && positionIndex < len(HangmanPositions) {
		fmt.Println(HangmanPositions[positionIndex])
	} else {
		// Si l'index est en dehors des limites, affiche un message d'erreur
		fmt.Println("Aucune position disponible pour cette tentative.")
	}
}
