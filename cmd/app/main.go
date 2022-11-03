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
		err := entries.Save(conf)
		utility.CheckErrors(err)
	}
}
