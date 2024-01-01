package hangman

func IsLetter(s string) bool {
	return len(s) == 1 && s >= "a" && s <= "z" || s >= "A" && s <= "Z"
}

func LetterSearch(hangman *HanGman, input string) (bool, string) {
	hangman.Result = false
	test := []rune(hangman.Words)
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
