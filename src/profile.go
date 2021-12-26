package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//the structure of a profile
type Profile struct {
	Name        string
	Mains       []Character
	Secondaries []Character
	Pockets     []Character
	SwitchFC    string
	Region      string
	Notes       string
}

//gets as many characters as the user wishes, until they write done
func GetProfileCharacters(cType string) []Character {
	c := 0
	var cList []Character

	for c < 1 {
		fmt.Println("What are your " + cType + "? Type \"done\" when you are done.")
		input := GetUserInput()
		if strings.ToLower(input) == "done" {
			c += 1
		} else {
			cList = append(cList, *MatchCharacter(input, GetListOfCharacters()))
		}

	}
	return cList
}

//just prints a nice format for a profile
func ProfilePrettyPrint(p Profile) string {
	mainNames := CharactersToString(p.Mains)
	secNames := CharactersToString(p.Secondaries)
	pocNames := CharactersToString(p.Pockets)

	ppp := `Profile Information of ` + p.Name + `
Mains: ` + mainNames + `
Secondaries: ` + secNames + `
Pockets: ` + pocNames + `
Switch FC: ` + p.SwitchFC + `
Region: ` + p.Region + `
Notes: ` + p.Notes

	return ppp
}

//asks the user for profile information to store
func GetNewProfileInfo() Profile {
	profile := Profile{}

	fmt.Println("What is the name of the profile?")
	profile.Name = GetUserInput()

	mList := GetProfileCharacters("mains")
	profile.Mains = mList

	sList := GetProfileCharacters("secondaries")
	profile.Secondaries = sList

	pList := GetProfileCharacters("pockets")
	profile.Pockets = pList

	fmt.Println("What is your Switch Friend Code?")
	profile.SwitchFC = GetUserInput()

	fmt.Println("What is your Region?")
	profile.Region = GetUserInput()

	fmt.Println("Any special notes?")
	profile.Notes = GetUserInput()

	fmt.Println("Profile saved! \n" + ProfilePrettyPrint(profile))

	return profile

}

//saves a profile in the according file. one in each file.
func SaveProfile(p Profile) {
	file, err := json.MarshalIndent(p, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("./profiles/"+p.Name+".json", file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//searches a profile in the according folder
func SearchProfile(inp string) Profile {
	file, err := ioutil.ReadFile("./profiles/" + inp + ".json")
	if err != nil {
		log.Fatal(err)
	}

	var prof Profile

	err = json.Unmarshal(file, &prof)
	if err != nil {
		log.Fatal(err)
	}

	return prof
}

//deletes a profile file
func DeleteProfile(inp string) {
	err := os.Remove("./profiles/" + inp + ".json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully deleted the profile of " + inp)
}
