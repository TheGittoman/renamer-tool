package filter

import (
	"encoding/json"
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
	filters := getFilters(conf)
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	persons := []person.Person{}
	found := false
	// way to delete stuff more efficiently in this
	// stage we have person object and we should filter it for incorrect instances
	for _, v := range names {
		p, _ := person.New(v, filters.Symbols)
		// if there is less than two names and the name lenght is less than 2 letters
		// or there is more than 4 names then skip
		// if ((len(p.Names) < 2 && len(p.Names[0]) <= 3) || len(p.Names) > 4) || len(p.Names[0]) > 15 {
		// 	continue
		// }
		// if len(p.Names) == 1 && len(p.Names[0]) < 4 {
		// 	continue
		// }
		// if len(p.Names) == 1 && len(p.Names[0]) > 15 {
		// 	continue
		// }
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
func WalkPath(conf *config.Config) ([]string, error) {
	oldNames := []string{}
	names := []string{}
	maxLenght := 0
	err := filepath.Walk(conf.ScanFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		// split path /path/path/interesting-path with \\ filter
		if info.IsDir() {
			splitPath := strings.Split(path, "\\")
			if maxLenght < len(splitPath) {
				oldNames = splitPath
				maxLenght = len(splitPath)
				return nil
			}
			// if length of the path is same as lenght of the split path
			// and temporary path not nil add last tip of the path to the list
			if maxLenght == len(splitPath) && oldNames != nil {
				names = append(names, oldNames[len(oldNames)-1])
				oldNames = nil
			}
			names = append(names, splitPath[len(splitPath)-1])
		}
		return nil
	})
	return names, err
}
