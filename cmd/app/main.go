package main

import (
	"fmt"
	"renamer-tool/src/config"
	"renamer-tool/src/filter"
	"renamer-tool/src/person"
	"renamer-tool/src/utility"
)

func main() {
	runSearch := false
	conf := config.New()

	if runSearch {
		fmt.Println("Searching for files")
		// try to open a filepath for scrawling
		names, paths, err := filter.WalkPath(conf)
		utility.CheckErrors(err)
		persons := filter.Result(names, conf)
		person.SaveToFile(persons, paths, conf.SaveFolder)
		fmt.Printf("Operation successful: %d entries added", len(persons))
	}
	if !runSearch {
		fmt.Println("running without search")
		entries := person.ReadFromFile(conf.SaveFolder)
		for i := 0; i < 20; i++ {
			fmt.Println(entries.Entries[i].Names)
		}
		count := 0
		for name, path := range entries.Paths {
			count++
			if count == 20 {
				break
			}
			fmt.Println(name + "\n" + path)
		}
		err := entries.Save()
		utility.CheckErrors(err)
	}
}
