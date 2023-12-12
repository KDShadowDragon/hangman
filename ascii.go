package hangman

import (
	"bufio"
	"os"
)

func AsciiDraw(filename2 string, str string) string {
	art := ""
	data := [][]string{}
	line := 1
	for e := range str {
		data = append(data, []string{})
		line = 1
		file, _ := os.Open(filename2)
		fileScanner := bufio.NewScanner(file)
		for fileScanner.Scan() {
			if fileScanner.Text() == "" || fileScanner.Text() == " " {
				line++
			} else if line == int(rune(str[e]))-30 {
				data[e] = append(data[e], fileScanner.Text())
			}
		}
		file.Close()
	}
	for i := 0; i < 8; i++ {
		for e := range data {
			art += data[e][i]
		}
		art += "\n"
	}
	return art
}

func HangmanArt(hangman *HanGman) string {
	filename, _ := os.Open("hangman-Art/hangman.txt")
	fileScanner := bufio.NewScanner(filename)
	line := 1
	result := ""
	if hangman.Attemps == 0 {
		return result
	} else {
		for fileScanner.Scan() {
			if fileScanner.Text() == "" || fileScanner.Text() == " " {
				line++
			} else if line == hangman.Attemps {
				result += "\n"
				result += fileScanner.Text()
			}
		}
		return result
	}
}
