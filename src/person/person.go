package person

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"renamer-tool/src/utility"
	"strings"
)

// Person class
type Person struct {
	Names []string `json:"Names"`
}

// Actresses class with all person members
type Actresses struct {
	ListOfActresses []Person `json:"Actresses"`
}

// New create new member of actor class
func New(input string, runes []string) (Person, error) {
	if len(input) == 0 || len(runes) == 0 {
		err := errors.New("input string is empty")
		log.Fatalln(err)
	}
	var names, oldNames []string
	if len(input) > 0 {
		names = strings.Split(input, string(runes[len(runes)-1]))
	}
	for _, iRune := range runes {
		oldNames = strings.Split(input, iRune)
		if len(names) < len(oldNames) {
			names = oldNames
		}
	}
	p := Person{names}
	return p, nil
}

// SaveToFile does what it says
func SaveToFile(actorList []Person, filename string) {
	actresses := new(Actresses)
	actresses.ListOfActresses = actorList
	file, _ := json.MarshalIndent(actresses, "", " ")
	ioutil.WriteFile(filename, file, 0755)
	// comment comment comment
}

// ReadFromFile function for reading actresses from generated json file
func ReadFromFile(filename string) *Actresses {
	file, err := ioutil.ReadFile(filename)
	utility.CheckErrors(err)
	actresses := new(Actresses)
	err = json.Unmarshal(file, actresses)
	utility.CheckErrors(err)
	return actresses
}
