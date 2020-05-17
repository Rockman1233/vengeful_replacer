package config

import (
	"errors"
	"vengeful_replacer/app/algorithms"
	"vengeful_replacer/app/dictionaries"
	"vengeful_replacer/app/scanners"
)

type Config struct {
	Dictionary dictionaries.Dictionary
	Algorithm  algorithms.Algorithm
	Scanner    scanners.Scanner
}

func New(params ...interface{}) (Config, error) {

	newEntity := Config{}

	for _, val := range params {
		switch entity := val.(type) {
		case dictionaries.Dictionary:
			newEntity.Dictionary = entity
			break
		case scanners.Scanner:
			newEntity.Scanner = entity
			break
		case algorithms.Algorithm:
			newEntity.Algorithm = entity
			break
		default:
			return newEntity, errors.New("half-full config")
		}
	}

	return newEntity, nil

}
