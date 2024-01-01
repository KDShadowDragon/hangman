package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func RandomWord(filename string) (string, string) {
	var word string
	var words string
	var stop error
	word, stop = loadWordList(filename)
	fmt.Println("|Continue...| Random Word Choose Sucessfuly")
	if stop != nil {
		fmt.Println("|Error 002| Error with the Randomizer of the Word")
		os.Exit(1)
	}
	for len(word) != len(words) {
		words += "_"
	}
	return word, words
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
