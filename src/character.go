package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

//the list of characters in the json file
type CharacterList struct {
	Character []Character `json:"Characters"`
}

//structure of the characters in the json file
type Character struct {
	Name      string   `json:"name"`
	Id        string   `json:"id"`
	Series    string   `json:"series"`
	FirstGame string   `json:"first-game"`
	Asset     string   `json:"asset"`
	Aliases   []string `json:"aliases"`
}

func GetListOfCharacters() CharacterList {
	file, err := ioutil.ReadFile("./data/characters.json")
	if err != nil {
		log.Fatal(err)
	}

	var char CharacterList

	err = json.Unmarshal(file, &char)
	if err != nil {
		log.Fatal(err)
	}

	return char
}

//gets the character details in a neat format
func CharacterDetails(char Character) string {
	return `Your character is: **` + char.Name + `**
ID: ` + char.Id + `
From the ` + char.Series + ` Series
First Smash game appearance: ` + char.FirstGame + `
Asset: ` + char.Asset + `
Aliases: ` + fmt.Sprint(strings.Join(char.Aliases, ", "))
}

//returns a readable string from a list of characters
func CharactersToString(chars []Character) string {
	var cList []string

	for _, c := range chars {
		cList = append(cList, c.Name)
	}

	return strings.Join(cList, ", ")
}

//returns a pointer to a character when it finds a match
func MatchCharacter(input string, charList CharacterList) *Character {
	//first looks for an exact match
	for i := 0; i < CharacterCount; i++ {
		if charList.Character[i].Name == input || charList.Character[i].Id == input {
			char := charList.Character[i]
			return &char
		} else {
			//lastly it looks through the aliases to find a match
			for _, name := range charList.Character[i].Aliases {
				if name == input {
					char := charList.Character[i]
					return &char
				}
			}
		}
	}
	//if there really is no match found, returns nil
	return nil
}

//same as above, but searches for multiple characters, maybe a lookup per series or game, or just in case of echoes (paisy)
func MatchMultipleCharacters(input string, charList CharacterList) []*Character {
	var matches []*Character

	for i := 0; i < CharacterCount; i++ {
		if charList.Character[i].Name == input || charList.Character[i].Id == input || charList.Character[i].FirstGame == input || charList.Character[i].Series == input {
			matches = append(matches, &charList.Character[i])

		} else {
			for _, name := range charList.Character[i].Aliases {
				if name == input {
					matches = append(matches, &charList.Character[i])

				}
			}
		}
	}
	//returns the list of matches, may be empty
	return matches

}
