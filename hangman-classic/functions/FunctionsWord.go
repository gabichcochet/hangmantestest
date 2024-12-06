package Hangman

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

var words []string

// Addwords lit des mots à partir du fichier words.txt et les ajoute au slice words.
func Addwords(filepath string) {
	// Utilisation d'un chemin relatif vers le fichier words.txt
	 // Chemin relatif correct depuis Hangman-Web

	// Ouvre le fichier
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Lit chaque ligne du fichier et ajoute chaque mot au slice words
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
// GetRandomWord retourne un mot aléatoire du slice words.
func GetRandomWord() string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(words))
	return words[randomIndex]
}


// UpdateWord met à jour l'état actuel du mot avec la lettre devinée si elle existe dans le mot.
func UpdateWord(word string, currentWordState []string, letter string) bool {
	found := false
	for i, char := range word {
		if string(char) == letter {
			currentWordState[i] = letter
			found = true
		}
	}
	return found
}