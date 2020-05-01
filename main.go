package main

import (
	"fmt"
	"os"
	"vengeful_replacer/app/scanner"
)

func main() {

	fmt.Print("insert y value here: ")

	newEntityOfScanner, _ := scanner.New(*os.Stdin)
	fmt.Println(newEntityOfScanner.GetData())
}
