package app

import "vengeful_replacer/app/scanners"

type Application struct {
	scanner scanners.Scanner
	//algorithm
	//output
}

func New() (Application, error) {

	return Application{}, nil
}
