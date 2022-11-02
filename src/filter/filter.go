package filter

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	actorList := []person.Person{}
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
			actorList = append(actorList, p)
		}
		foundNumbers = 0
		found = false
	}
	return actorList
}
