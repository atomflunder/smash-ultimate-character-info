package src

import (
	"fmt"
	"log"
)

func MainMenu() {
	fmt.Println(`Welcome to the Smash Ultimate Character Info Menu!
What do you wanna do?
1) Create, view or delete profiles
2) Search for characters`)

	ui := GetUserInput()

	if ui == "1" {
		ProfileMenu()
	} else if ui == "2" {
		SearchMenu()
	} else {
		fmt.Println("Please choose a valid input.")
		return
	}
}

func ProfileMenu() {
	fmt.Println(`Welcome to the Profile Menu!
What do you wanna do?
1) Create a new profile
2) View a profile
3) Delete a profile`)
	ui := GetUserInput()

	if ui == "1" {
		SaveProfile(GetNewProfileInfo())
	} else if ui == "2" {
		fmt.Println("Which profile are you looking for?")
		input := GetUserInput()
		prof := SearchProfile(input)
		fmt.Println(ProfilePrettyPrint(prof))
	} else if ui == "3" {
		fmt.Println("Which profile do you want to delete?")
		input := GetUserInput()
		DeleteProfile(input)
	} else {
		fmt.Println("Please choose a valid input.")
	}
}

func SearchMenu() {
	fmt.Println(`Welcome to the Search Menu!
What do you wanna do?
1) Pick a random character
2) Look up a character
3) Search for multiple characters at once
...
9) Help for searching`)

	ui := GetUserInput()

	if ui == "1" {
		charList := GetListOfCharacters()
		randChar := RandomCharacter(charList)
		fmt.Println(CharacterDetails(randChar))

	} else if ui == "2" {
		fmt.Println("Which Character do you want to look up? You can search for Names or Fighter Numbers.")

		inp := GetUserInput()

		charList := GetListOfCharacters()
		char := MatchCharacter(inp, charList)
		if char == nil {
			log.Fatal("Could not find this character.")
		}
		//matchcharacter returns a pointer to a character, so we need to use that here
		fmt.Println(CharacterDetails(*char))
	} else if ui == "3" {
		fmt.Println("Which characters do you want to look up? You can search for Names, Fighter Numbers, Origin Series or Smash Game appearances.")

		inp := GetUserInput()

		charList := GetListOfCharacters()
		chars := MatchMultipleCharacters(inp, charList)

		if len(chars) == 0 {
			log.Fatal("No characters found.")
		} else {
			for _, c := range chars {
				fmt.Println(CharacterDetails(*c) + "\n")
			}
		}

	} else if ui == "9" {
		fmt.Println(`
Welcome to the help menu. If you cannot find a character, here is what you need to know:

All of the fighters use the official name as they appear in game. However you can also search them by common nicknames. 
Example: "Zero Suit Samus" and "ZSS" will both work.

Capitalisation does not matter, every character is saved in all lowercase and your input will get converted automatically.

You can also search Fighters by their official fighter number as they appear on the smashbros.com website. Echoes are marked with an "e" behind the number. 
Example: 4 -> Samus, 4e -> Dark Samus.

If you search for multiple characters, you can also search them by their Origin Series (as listed on the smashbros.com website).
Example: Searching "Metroid" will return Samus, Dark Samus, Zero Suit Samus and Ridley.

You can also search multiple characters by their first Smash Game appearance, their titles are abbreviated however. Here they are listed:
64, Melee, Brawl, Wii U, Ultimate
Example: Searching "Melee" will return the 14 characters who gave their debut in Melee. From Peach to Mr. Game & Watch

If you only search for one character, but use a nickname multiple characters have in common, the first in order of fighters numbers will get returned. 
If you search for multiple, both will be returned.
Example: If you search for one character with "Paisy", you will get Peach. If you search for multiple with the same input you will get both Peach and Daisy.

Please keep in mind that sometimes 1 playable character will be split up into multiple ones, according to the smashbros.com website.
Example: Pokémon Trainer is split up into Squirtle, Ivysaur and Charizard.

Hope this was useful!`)
	} else {
		fmt.Println("Please choose a valid input.")
		return
	}
}
