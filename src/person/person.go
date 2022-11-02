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
	Entries []Person `json:"Entries"`
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
func SaveToFile(persons []Person, filename string) {
	entries := new(Entries)
	entries.Entries = persons
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

/*
write to file from folder structure and try to delete trash results
read from file and add all to entries class
search images with name including names in json result from the first step

	trash-actressfirstname-actresslastname-trash.jpg > actressfirstname-actresslastname.jpg

make new path and save /folder/<actress>/<actress>-01.jpg

way to add to the json when running instead of deleting everything
*/
