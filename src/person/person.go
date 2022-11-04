package person

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"renamer-tool/src/config"
	"renamer-tool/src/utility"
	"strings"

	"github.com/Kagami/go-face"
)

// Person class
type Person struct {
	Names []string `json:"Names"`
	// Identifiers []face.Face `json:"Identifiers`
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

// ReadNames returns all names of person divided by a symbol
func (e *Person) ReadNames(symbol string) string {
	if e == nil {
		log.Fatal("ReadNames failes because person was nil")
	}
	fileName := ""
	for index, name := range e.Names {
		if index == 0 {
			fileName += name
			continue
		}
		fileName += symbol + name
	}
	return fileName
}

// Save and format the entries with picture paths
func (e *Entries) Save(conf *config.Config) error {
	separator := "_"
	found := false
	count := 0
	for _, actor := range e.Entries {
		for fileName, path := range e.Paths {
			for _, name := range actor.Names {
				if !strings.Contains(fileName, name) {
					found = false
					break
				}
				found = true
			}
			if !found {
				continue
			}
			if !FindFace(path) {
				fmt.Println("can't find a face")
				continue
			}
			count++
			fmt.Println("Saving to file!")
			saveDir := actor.ReadNames("_")
			file, err := ioutil.ReadFile(path)
			utility.CheckErrors(err)
			if count > 1 {
				err := ioutil.WriteFile("./data/tagged_pictures/"+saveDir+separator+fmt.Sprint(count)+".jpg", file, 755)
				utility.CheckErrors(err)
			} else {
				err := ioutil.WriteFile("./data/tagged_pictures/"+saveDir+".jpg", file, 755)
				utility.CheckErrors(err)
			}
		}
		count = 0
	}
	return nil
}

// FindFace returns true if given image/path image includes a face
func FindFace(path string) bool {
	rec, err := face.NewRecognizer("F:/Git/Coding/dlib-models")
	if err != nil {
		log.Fatalf("Can't init face recognizer: %v", err)
	}
	defer rec.Close()
	fmt.Println("Scanning faces")
	faces, err := rec.RecognizeFile(path)
	if err != nil {
		return false
	}
	if len(faces) > 0 {
		return true
	}
	return false
}

// FindFace returns true if given image/path image includes a face
// func FindFace(path *string) bool {
// 	rec := recognizer.Recognizer{}
// 	err := rec.Init("F:/Git/Coding/dlib-models")
// 	if err != nil {
// 		log.Fatal("can't find dlib-models")
// 		return false
// 	}
// 	rec.Tolerance = 0.4
// 	rec.UseGray = true
// 	rec.UseCNN = false
// 	defer rec.Close()

// 	fmt.Println("scanning for faces...")
// 	_, err = rec.RecognizeMultiples(*path)
// 	if err != nil {
// 		return false
// 	}

// 	return true
// }

/*
search images with name including names in json result from the first step

	trash-actressfirstname-actresslastname-trash.jpg > actressfirstname-actresslastname.jpg

make new path and save /folder/<actress>/<actress>-01.jpg

way to add to the json when running instead of deleting everything
*/
