package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"renamer-tool/src/person"
	"strings"
)

type filterWords struct {
	Words []string `json:"Words"`
}

func getFilters() []string {
	words := new(filterWords)
	file, err := ioutil.ReadFile("./data/filters.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, words)
	if err != nil {
		log.Fatal(err)
	}
	return words.Words
}

func main() {
	names := []string{}
	oldName := []string{}
	maxLenght := 0

	err := filepath.Walk("C:/Users/Kone/4chan/PornstarList/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		// split path /path/path/interesting-path with \\ filter
		if info.IsDir() {
			splitPath := strings.Split(path, "\\")
			if maxLenght < len(splitPath) {
				oldName = splitPath
				maxLenght = len(splitPath)
				return nil
			}
			// if length of the path is same as lenght of the split path
			// and temporary path not nil add last tip of the path to the list
			if maxLenght == len(splitPath) && oldName != nil {
				names = append(names, oldName[len(oldName)-1])
				oldName = nil
			}
			names = append(names, splitPath[len(splitPath)-1])
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	//filters from json that are used to filter the results
	filters := getFilters()
	actorList := []person.Person{}
	found := false
	// way to delete stuff more efficiently in this
	// stage we have person object and we should filter it for incorrect instances
	for _, v := range names {
		p, _ := person.New(v)
		// if there is less than two names and the name lenght is less than 2 letters
		// or there is more than 4 names then skip
		if ((len(p.Names) < 2 && len(p.Names[0]) <= 3) || len(p.Names) > 4) || len(p.Names[0]) > 15 {
			continue
		}
		// go through filters and for each filter test all the names
		for _, filter := range filters {
			for _, name := range p.Names {
				if strings.Contains(name, filter) {
					found = true
				}
			}
		}
		if !found {
			actorList = append(actorList, p)
		}
		found = false
	}

	person.SaveToFile(actorList, "./data/actresses.json")
	fmt.Printf("Operation successful: %d actresses added", len(actorList))
}
