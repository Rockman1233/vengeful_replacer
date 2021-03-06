package algorithms

import (
	"vengeful_replacer/app/dictionaries"
)

type EmptyAlgorithm struct {
	data       string
	result     string
	dictionary dictionaries.Dictionary
}

func (alg *EmptyAlgorithm) Run() {

	var result []rune
	runeString := []rune(alg.data)

	for _, letter := range runeString {

		replacer := changeChar(letter, alg.dictionary)

		if replacer == 0 {
			replacer = letter
		}

		result = append(result, replacer)

	}

	alg.result = string(result)
}

func changeChar(runeChar rune, dict dictionaries.Dictionary) rune {

	var result rune

	for origin, replacer := range dict.GetMap() {

		originRune := []rune(origin)[0]
		originReplacer := []rune(replacer)[0]

		if originRune == runeChar {
			result = originReplacer
		}
	}

	if &result != nil {
		return result
	}

	return 0
}

func (alg EmptyAlgorithm) GetResult() string {
	return alg.result
}

func (alg *EmptyAlgorithm) SetDictionary(dictionary dictionaries.Dictionary) {
	alg.dictionary = dictionary
}

func (alg *EmptyAlgorithm) SetData(data string) {
	alg.data = data
}

func NewEmptyAlgorithm() *EmptyAlgorithm {

	newEntity := EmptyAlgorithm{}

	return &newEntity
}
