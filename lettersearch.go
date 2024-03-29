package hangman

func IsLetter(s string) bool {
	return s >= "a" && s <= "z" || s >= "A" && s <= "Z"
}

func LetterSearch(hangman *HanGman, input string) (bool, string) {
	if !IsLetter(input) {
		hangman.Result = true
		return hangman.Result, ""
	}
	hangman.Result = false
	test := []rune(hangman.Words)
	compteur := hangman.Attemps - 1
	for compteur != 11 {
		if hangman.LooseLetter[compteur] == input {
			hangman.Result = true
			return hangman.Result, ""
		} else {
			compteur++
			hangman.Result = false
		}
	}
	if len(input) == 1 {
		for j := 0; j < len(hangman.Words); j++ {
			if string(hangman.Word[j]) == input {
				test[j] = rune(input[0])
				hangman.Result = true
				hangman.WinLetter[hangman.Attemps] = string(test)
			}
		}
	}
	if hangman.Result {
		hangman.Words = string(test)
	}
	if hangman.Word == input {
		input = "!"
	}
	return hangman.Result, input
}
