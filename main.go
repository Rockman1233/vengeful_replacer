package main

import (
	"fmt"
	"os"
	"vengeful_replacer/app"
	"vengeful_replacer/app/algorithms"
	"vengeful_replacer/app/config"
	"vengeful_replacer/app/dictionaries"
	"vengeful_replacer/app/scanners"
)

func main() {

	fmt.Print("insert value here: ")

	// Get data from stdin
	newEntityOfScanner := scanners.NewStdinScanner()
	err := newEntityOfScanner.SetSource(*os.Stdin)

	// Implement dictionary and algorithm
	newEntityOfDictionary := dictionaries.NewSimpleDictionary()
	newEntityOfAlgorithm := algorithms.NewEmptyAlgorithm()

	// Load data to config
	configEntity, err := config.New(newEntityOfScanner, newEntityOfAlgorithm, newEntityOfDictionary)

	// Load config to data
	appEntity := app.New(configEntity)

	// Run application
	appEntity.Run()

	if err != nil {
		panic(err)
	}

	// Get result
	fmt.Println(appEntity.GetResult())
}
