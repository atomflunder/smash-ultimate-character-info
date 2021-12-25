package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func MainMenu() {
	fmt.Println(`Welcome to the Smash Ultimate Character Info Tool!
What do you wanna do?
1) Pick a random character
2) Look up a character`)

	var ui int
	_, err := fmt.Scanln(&ui)
	if err != nil {
		log.Fatal(err)
	}

	if ui == 1 {
		charList := GetListOfCharacters()
		randChar := RandomCharacter(charList)
		fmt.Println(CharacterDetails(randChar))

	} else if ui == 2 {
		fmt.Println("Which Character do you want to look up?")
		inpReader := bufio.NewReader(os.Stdin)
		inp, _ := inpReader.ReadString('\n')

		//since the json file is in all lowercase, we convert the input
		inp = strings.ToLower(inp)
		//we have to trim the whitespace from the user input, too
		inp = strings.TrimSpace(inp)

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
