package Hangman


// InitializeWord crée un slice représentant l'état actuel du mot, rempli de underscores.
func InitializeWord(word string) []string {
	currentWordState := make([]string, len(word))
	for i := range currentWordState {
		currentWordState[i] = "_"
	}
	return currentWordState
}

// InitializeWord crée un slice représentant l'état actuel du mot, rempli de underscores.

// initialisation et boucle principale
// initialisation et boucle principale

// Fonction pour convertir la lettre en minuscule
