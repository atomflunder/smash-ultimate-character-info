package src

import (
	"fmt"
	"log"
)

func MainMenu() {
	fmt.Println(`Welcome to the Smash Ultimate Character Info Tool!
What do you wanna do?
1) Pick a random character
2) Look up a character`)

	ui := GetUserInput()

	if ui == "1" {
		charList := GetListOfCharacters()
		randChar := RandomCharacter(charList)
		fmt.Println(CharacterDetails(randChar))

	} else if ui == "2" {
		fmt.Println("Which Character do you want to look up?")

		inp := GetUserInput()

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
