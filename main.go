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

	fmt.Print("insert y value here: ")

	newEntityOfScanner := scanners.NewStdinScanner()
	err := newEntityOfScanner.SetSource(*os.Stdin)
	newEntityOfDictionary := dictionaries.NewSimpleDictionary()
	newEntityOfAlgorithm := algorithms.NewEmptyAlgorithm()

	configEntity, err := config.New(newEntityOfScanner, newEntityOfAlgorithm, newEntityOfDictionary)
	appEntity := app.New(configEntity)

	appEntity.Run()

	if err != nil {
		panic(err)
	}

	fmt.Println(appEntity.GetResult())
}
