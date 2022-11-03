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

// Entries class with all person members
type Entries struct {
	Entries []Person          `json:"Entries"`
	Paths   map[string]string `json:"Paths"`
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
func SaveToFile(persons []Person, paths map[string]string, filename string) {
	entries := new(Entries)
	entries.Entries = persons
	entries.Paths = paths
	file, _ := json.MarshalIndent(entries, "", " ")
	ioutil.WriteFile(filename, file, 0755)
	// comment comment comment
}

// ReadFromFile function for reading entries from generated json file
func ReadFromFile(filename string) *Entries {
	file, err := ioutil.ReadFile(filename)
	utility.CheckErrors(err)
	entries := new(Entries)
	err = json.Unmarshal(file, entries)
	utility.CheckErrors(err)
	return entries
}

// Save and format the entries with picture paths
func (e *Entries) Save() error {
	found := false
	count := 0
	for _, actor := range e.Entries {
		for fileName, path := range e.Paths {
			for _, name := range actor.Names {

			}
		}
	}
	count = 0
}

/*
search images with name including names in json result from the first step

	trash-actressfirstname-actresslastname-trash.jpg > actressfirstname-actresslastname.jpg

make new path and save /folder/<actress>/<actress>-01.jpg

way to add to the json when running instead of deleting everything
*/
