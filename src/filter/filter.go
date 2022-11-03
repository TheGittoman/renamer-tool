package filter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"renamer-tool/src/config"
	"renamer-tool/src/person"
	"strings"
)

type filterWords struct {
	Words   []string `json:"Words"`
	Symbols []string `json:"Symbols"`
}

func getFilters(conf *config.Config) filterWords {
	words := new(filterWords)
	file, err := ioutil.ReadFile(conf.FilterFolder)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, words)
	if err != nil {
		log.Fatal(err)
	}
	return *words
}

// Result filters names from list and returns them as a person object
func Result(names []string, conf *config.Config) []person.Person {
	//filters from json that are used to filter the results
	if len(names) == 0 {
		log.Fatalln("names list is missing")
	}
	// getFilters from conf.FilterFolder
	filters := getFilters(conf)
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	persons := []person.Person{}
	found := false
	// way to delete stuff more efficiently in this
	// stage we have person object and we should filter it for incorrect instances
	for _, v := range names {
		p, _ := person.New(v, filters.Symbols)
		// check the lenght of the list
		if len(p.Names) <= 1 {
			continue
		}
		// // go through filters and for each filter test all the names
		foundNumbers := 0
		for _, filter := range filters.Words {
			for _, name := range p.Names {
				if strings.Contains(name, filter) {
					found = true
					break
				}
				for _, number := range numbers {
					if strings.Contains(name, number) {
						foundNumbers++
						if foundNumbers > 4 {
							found = true
							break
						}
					}
				}
			}
		}
		if !found {
			persons = append(persons, p)
		}
		foundNumbers = 0
		found = false
	}
	return persons
}

// WalkPath path that will be scrawled by the command
// Takes config struct that included filedirs for commands
func WalkPath(conf *config.Config) ([]string, map[string]string, error) {
	names := []string{}
	pictures := make(map[string]string)
	fileTypes := []string{".png", ".jpg", ".jpeg"}
	err := filepath.Walk(conf.ScanFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fileName := strings.ToLower(info.Name())
		for _, fileType := range fileTypes {
			if strings.Contains(fileName, fileType) {
				file, err := ioutil.ReadFile(path)
				if err != nil {
					log.Println(err)
				}
				savePath := "./data/pictures/" + fileName
				if _, err := os.Stat(savePath); errors.Is(err, os.ErrNotExist) {
					err = ioutil.WriteFile(savePath, file, 755)
					if err != nil {
						log.Println(err)
					}
					pictures[fileName] = savePath
				}
			}
		}
		// get directory name
		if info.IsDir() {
			names = append(names, fileName)
		}
		return nil
	})
	return names, pictures, err
}

// kuvien ja nimien vertailu
// found = true if not found turn false and go to next picture
