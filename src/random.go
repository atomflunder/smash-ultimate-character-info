package src

import (
	"math/rand"
	"time"
)

//the total character count in smash, 82 normal fighters + 7 echoes
var CharacterCount int = 89

func RandomCharacter(charList CharacterList) Character {
	//makes sure to get a random character each time
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(CharacterCount)
	//gets a random character and returns it
	randChar := charList.Character[i]
	return randChar
}
