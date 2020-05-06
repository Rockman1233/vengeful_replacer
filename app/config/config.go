package config

import (
	"go/scanner"
	"vengeful_replacer/app/algorithms"
	"vengeful_replacer/app/dictionaries"
)

type Config struct {
	dict      dictionaries.Dictionary
	algorithm algorithms.Algorithm
	scanner   scanner.Scanner
}

func New(params ...interface{}) (Config, err) {

	newEntity := Config{}

	for _, val := range params {
		switch val.(type) {
		case dictionaries.Dictionary:
			newEntity.dict = val
			break
		case scanner.Scanner:
			newEntity.scanner = val
			break
		case algorithms.Algorithm:
			newEntity.algorithm = val
			break
		default:

		}
	}

	return newEntity, err

}
