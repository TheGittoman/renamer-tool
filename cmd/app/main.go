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
	filters := getFilters()
	names := []string{}
	maxLenght := 0
	oldName := []string{}
	actorList := []person.Person{}
	err := filepath.Walk("C:/Users/Kone/4chan/PornstarList/www.freeones.com", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if info.IsDir() {
			for _, v := range filters {
				if strings.Contains(path, v) {
					return nil
				}
			}
			splitPath := strings.Split(path, "\\")
			if maxLenght < len(splitPath) {
				oldName = splitPath
				maxLenght = len(splitPath)
				return nil
			}
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
	for _, v := range names {
		p, err := person.New(v)
		if err != nil {
			fmt.Printf("p.Names: %v\n : %s\n", p.Names, err)
		}
		actorList = append(actorList, p)
	}
	person.SaveToFile(actorList, "./data/actresses.json")
	fmt.Printf("Operation successful: %d actresses added", len(actorList))
}
