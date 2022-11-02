package main

import (
	"fmt"
	"os"
	"path/filepath"
	"renamer-tool/src/config"
	"renamer-tool/src/filter"
	"renamer-tool/src/person"
	"strings"
)

func main() {
	conf := config.New()
	names := []string{}
	oldName := []string{}
	maxLenght := 0

	fmt.Printf("Filters folder set to: 	%s \nSave folder set to: 	%s\nScan folder set to: 	%s\n",
		conf.FilterFolder, conf.SaveFolder, conf.ScanFolder)
	err := filepath.Walk(conf.ScanFolder, func(path string, info os.FileInfo, err error) error {
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

	actorList := filter.Result(names, conf)
	person.SaveToFile(actorList, conf.SaveFolder)
	fmt.Printf("Operation successful: %d entries added", len(actorList))
}
