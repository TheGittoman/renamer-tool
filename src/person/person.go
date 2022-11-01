package person

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

// Person class
type Person struct {
	Names []string `json:"Names"`
}

type Actresses struct {
	ListOfActresses []Person `json:"Actresses"`
}

// New create new member of actor class
func New(input string, runes []string) (Person, error) {
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

// ReadFromFile function
func ReadFromFile(filename string) []Person {
	file, err := ioutil.ReadFile(filename)
}
