package hangman

import (
	"encoding/json"
	"fmt"
	"os"
)

func Save(hangman *HanGman) {
	var save string
	Datahangman, err := json.Marshal(*hangman)
	if CheckErrSave(err) {
		fmt.Print("create a name for your save: ")
		fmt.Scan(&save)
		save = "Save/" + save + ".txt"
		os.WriteFile(save, Datahangman, 0644)
		fmt.Println("game saved in", save)
		os.Exit(0)
	}
}

func CheckErrSave(err error) bool {
	status := true
	if err != nil {
		fmt.Println("A error have been seen when the saving of the game")
		status = false
		return status
	} else {
		return status
	}
}
