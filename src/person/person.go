package person

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

// Person class
type Person struct {
	Names []string `json:"Names"`
}

// New create new member of actor class
func New(name string) (Person, error) {
	names := strings.Split(name, "-")
	if len(names) >= 2 {
		p := Person{names}
		return p, nil
	}
	names = strings.Split(names[0], "_")
	if len(names) >= 2 {
		p := Person{names}
		return p, nil
	}
	p := Person{[]string{name}}
	return p, errors.New("Name is shorter than expected")
}

// SaveToFile does what it says
func SaveToFile(actorList []Person, filename string) {
	file, _ := json.MarshalIndent(actorList, "", " ")
	ioutil.WriteFile(filename, file, 0755)
}
