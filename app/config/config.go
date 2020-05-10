package config

import (
	"errors"
	"go/scanner"
	"vengeful_replacer/app/algorithms"
	"vengeful_replacer/app/dictionaries"
)

type Config struct {
	dict      dictionaries.Dictionary
	algorithm algorithms.Algorithm
	scanner   scanner.Scanner
}

func New(params ...interface{}) (Config, error) {

	newEntity := Config{}

	for _, val := range params {
		switch entity := val.(type) {
		case dictionaries.Dictionary:
			newEntity.dict = entity
			break
		case scanner.Scanner:
			newEntity.scanner = entity
			break
		case algorithms.Algorithm:
			newEntity.algorithm = entity
			break
		default:
			return newEntity, errors.New("half-full config")
		}
	}

	return newEntity, nil

}
