package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func RandomWord(filename string) string {
	var word string
	var stop error
	arg_len := len(os.Args[2:])
	if arg_len < 1 {
		fmt.Println("|Error 001| Not Enought Argument")
		os.Exit(1)
	} else if arg_len > 1 {
		fmt.Println("|Error 001| Too Much Argument")
		os.Exit(1)
	} else {
		fmt.Println("|Continue...| Search for a Random Word")
		word, stop = loadWordList(filename)
		fmt.Println("|Continue...| Random Word Choose Sucessfuly")
		if stop != nil {
			fmt.Println("|Error 002| Error with the Randomizer of the Word")
			os.Exit(1)
		}
	}
	return word
}

func loadWordList(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("|Error 002| Error with the name of your dictionnary")
		os.Exit(1)
	}
	defer file.Close()
	var wordList string
	var i int
	r := rand.Intn(34)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i++
		wordList = scanner.Text()
		if i == r {
			break
		}
	}
	return wordList, err
}
