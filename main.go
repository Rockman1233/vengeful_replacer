package main

import (
	"fmt"
	"os"
	"vengeful_replacer/app/scanners"
)

func main() {

	fmt.Print("insert y value here: ")

	newEntityOfScanner, _ := scanners.New(*os.Stdin)
	fmt.Println(newEntityOfScanner.GetData())
}
