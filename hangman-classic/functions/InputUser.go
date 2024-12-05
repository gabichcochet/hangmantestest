package Hangman

import (
	"bufio"
	"os"
)

// InputUser lit une ligne d'entrée de l'utilisateur.
func InputUser() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// InputUser lit une ligne d'entrée de l'utilisateur.
