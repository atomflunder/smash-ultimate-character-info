package src

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//the list of characters in the json file
type CharacterList struct {
	Character []Character `json:"Characters"`
}

//structure of the characters in the json file
type Character struct {
	Name       string   `json:"name"`
	Id         string   `json:"id"`
	Series     string   `json:"series"`
	First_game string   `json:"first-game"`
	Aliases    []string `json:"aliases"`
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
