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
	err := filepath.Walk("C:/Users/Kone/4chan/PornstarList/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if info.IsDir() {
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
	found := false
	for _, v := range names {
		p, _ := person.New(v)
		for _, v := range filters {
			for _, v2 := range p.Names {
				if strings.Contains(v2, v) {
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
