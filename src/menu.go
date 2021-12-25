package src

import (
	"fmt"
	"log"
	"strings"
)

func MainMenu() {
	fmt.Println(`Welcome to the Smash Ultimate Character Info Tool!
What do you wanna do?
1) Pick a random character
2) Look up a character`)

	var ui int
	_, err := fmt.Scan(&ui)
	if err != nil {
		log.Fatal(err)
	}

	if ui == 1 {
		charList := GetListOfCharacters()
		randChar := RandomCharacter(charList)
		fmt.Println(CharacterDetails(randChar))

	} else if ui == 2 {
		fmt.Println("Which Character do you want to look up?")
		var inp string
		_, err := fmt.Scan(&inp)
		if err != nil {
			log.Fatal(err)
		}

		//since the json file is in all lowercase, we convert the input
		inp = strings.ToLower(inp)

		charList := GetListOfCharacters()
		char := MatchCharacter(inp, charList)
		if char == nil {
			log.Fatal("Could not find this character.")
		}
		//matchcharacter returns a pointer to a character, so we need to use that here
		fmt.Println(CharacterDetails(*char))
	} else {
		fmt.Println("Please choose a valid input")
		return
	}
}
