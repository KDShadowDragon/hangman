package hangman

import (
	"fmt"
)

func IsLetter(s string) bool {
	return len(s) == 1 && s >= "a" && s <= "z" || s >= "A" && s <= "Z"
}

func LetterSearch(hangman *HanGman) (bool, string) {
	var wordtest string
	fmt.Scan(&wordtest)
	hangman.Result = false
	test := []rune(hangman.Words)
	if len(wordtest) == 1 {
		for j := 0; j < len(hangman.Words); j++ {
			if string(hangman.Word[j]) == wordtest {
				test[j] = rune(wordtest[0])
				hangman.Result = true
				hangman.WinLetter[hangman.Attemps] = string(test)
			}
		}
	}
	if hangman.Result {
		hangman.Words = string(test)
	}
	if hangman.Word == wordtest {
		wordtest = "!"
	}
	return hangman.Result, wordtest
}
