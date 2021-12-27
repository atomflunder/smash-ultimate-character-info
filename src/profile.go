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
	Name        string      `json:"Name"`
	Mains       []Character `json:"Mains"`
	Secondaries []Character `json:"Secondaries"`
	Pockets     []Character `json:"Pockets"`
	SwitchFC    string      `json:"SwitchFC"`
	Region      string      `json:"Region"`
	Notes       string      `json:"Notes"`
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
			char := MatchCharacter(input, GetListOfCharacters())
			if char == nil {
				fmt.Println("Could not find this character, please try again.")
			} else {
				cList = append(cList, *char)
			}

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

//searches a profile in the according folder, just get them by filename. can only return one profile
func SearchProfileByName(inp string) Profile {
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

//searching through all profiles by their main(s), returns multiple profiles
func SearchProfileByMain(inp string, pl []Profile) []Profile {
	var profList []Profile

	for _, p := range pl {
		for _, m := range p.Mains {
			if m.Name == strings.ToLower(inp) {
				profList = append(profList, p)
			}

		}
	}

	return profList
}

//searches all profiles by any character listed, returns multiple profiles
func SearchProfileByAnyCharacter(inp string, pl []Profile) []Profile {
	var profList []Profile

	for _, p := range pl {
		for _, m := range p.Mains {
			if m.Name == strings.ToLower(inp) {
				profList = append(profList, p)
			}
		}
		for _, s := range p.Secondaries {
			if s.Name == strings.ToLower(inp) {
				profList = append(profList, p)
			}
		}
		for _, po := range p.Pockets {
			if po.Name == strings.ToLower(inp) {
				profList = append(profList, p)
			}
		}
	}

	return profList
}

//searches all profiles by region, returns multiple profiles
func SearchProfileByRegion(inp string, pl []Profile) []Profile {
	var profList []Profile

	for _, p := range pl {
		if p.Region == strings.ToLower(inp) {
			profList = append(profList, p)
		}
	}

	return profList
}

//opens all profiles
func OpenAllProfiles() []Profile {
	var profList []Profile

	files, err := ioutil.ReadDir("./profiles/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		//we wanna skip the .gitignore and other files that may also be present in this directory
		if strings.HasSuffix(f.Name(), ".json") {
			file, err := ioutil.ReadFile("./profiles/" + f.Name())
			if err != nil {
				log.Fatal(err)
			}
			var prof Profile

			err = json.Unmarshal(file, &prof)
			if err != nil {
				log.Fatal(err)
			}

			profList = append(profList, prof)
		}
	}

	return profList
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

//deletes a profile file
func DeleteProfile(inp string) {
	err := os.Remove("./profiles/" + inp + ".json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully deleted the profile of " + inp)
}
