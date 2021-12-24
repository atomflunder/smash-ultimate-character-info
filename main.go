package main

import (
	"fmt"

	"github.com/phxenix-w/smash-ultimate-character-info/src"
)

func main() {
	fmt.Println("Welcome to the Smash Ultimate Character Info Tool!")

	fmt.Println("Picking a random character for you...")

	charList := src.GetListOfCharacters()

	//gets a random character
	var randChar = src.RandomCharacter(charList)

	//prints the character out in a neat format
	fmt.Println("Your character is: **" + randChar.Name + "**")
	fmt.Println("ID: " + randChar.Id)
	fmt.Println("From the " + randChar.Series + " Series")
	fmt.Println("First Smash game appearance: " + randChar.First_game)
	fmt.Println("Aliases: " + fmt.Sprint(randChar.Aliases))

}
