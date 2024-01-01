package hangman

import (
	"fmt"
	"os"
)

type HanGman struct {
	Replay      bool
	Result      bool
	Attemps     int
	Letter      string
	Words       string
	Word        string
	LooseLetter [16]string
	WinLetter   [16]string
}

func Hangman(filename string, filename2 string) {
	var hangman HanGman
	hangman.Word, hangman.Words = RandomWord(filename)
	for i, _ := range hangman.Word {
		i++
		hangman.Words += "_"
	}
	for hangman.Attemps = 0; hangman.Attemps != 11; {
		fmt.Println("\n", HangmanArt(&hangman))
		fmt.Println(AsciiDraw(filename2, hangman.Words))
		fmt.Println("Letter Already Choose : ", hangman.LooseLetter)
		fmt.Println("Attemps Left : ", hangman.Attemps, " Max 10 attemps")
		fmt.Print("Choose a Letter: ")
		hangman.Result, hangman.LooseLetter[hangman.Attemps] = LetterSearch(&hangman, "input")
		if hangman.LooseLetter[hangman.Attemps] == "Stop" || hangman.LooseLetter[hangman.Attemps] == "stop" {
			hangman.LooseLetter[hangman.Attemps] = ""
			Save(&hangman)
		}
		if len(hangman.LooseLetter[hangman.Attemps]) > 1 {
			hangman.Attemps++
			ClearScreen()
		}
		if !hangman.Result && hangman.Attemps == 0 {
			hangman.Attemps++
			ClearScreen()
		} else if hangman.Result && hangman.Attemps == 0 {
			if hangman.Word == hangman.Words {
				ClearScreen()
				fmt.Println(AsciiDraw(filename2, "GG You Found The Word : "), AsciiDraw(filename2, hangman.Word))
				fmt.Println("")
				fmt.Println("")
				fmt.Println("Would you want to restart ?")
				fmt.Println("true = Yes | else = No")
				_, err := fmt.Scan(&hangman.Replay)
				if err != nil {
					fmt.Println("|Error 003| The input have a bug please reload the program")
					os.Exit(1)
				} else if hangman.Replay {
					Hangman(filename, filename2)
				} else {
					ClearScreen()
					break
				}
			} else {
				ClearScreen()
				continue
			}
		} else if !hangman.Result && hangman.LooseLetter[hangman.Attemps] != hangman.LooseLetter[hangman.Attemps-1] {
			hangman.Attemps++
			ClearScreen()
		}
		if hangman.LooseLetter[hangman.Attemps] == hangman.LooseLetter[hangman.Attemps-1] {
			ClearScreen()
			continue

		}
		if hangman.LooseLetter[hangman.Attemps] == "!" || hangman.Word == hangman.Words {
			ClearScreen()
			fmt.Println(AsciiDraw(filename2, "GG You Found The Word : "), AsciiDraw(filename2, hangman.Word))
			fmt.Println("")
			fmt.Println("")
			fmt.Println("Would you want to restart ?")
			fmt.Println("true = Yes | else = No")
			_, err := fmt.Scan(&hangman.Replay)
			if err != nil {
				fmt.Println("|Error 003| The input have a bug please reload the program")
				os.Exit(1)
			} else if hangman.Replay {
				Hangman(filename, filename2)
			} else {
				ClearScreen()
				break
			}
		}
	}
}
