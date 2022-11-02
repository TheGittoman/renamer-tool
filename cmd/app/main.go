package main

import (
	"fmt"
	"renamer-tool/src/config"
	"renamer-tool/src/filter"
	"renamer-tool/src/person"
	"renamer-tool/src/utility"
)

func main() {
	conf := config.New()

	// try to open a filepath for scrawling
	names, err := filter.WalkPath(conf)
	utility.CheckErrors(err)

	persons := filter.Result(names, conf)
	person.SaveToFile(persons, conf.SaveFolder)
	fmt.Printf("Operation successful: %d entries added", len(persons))
}
