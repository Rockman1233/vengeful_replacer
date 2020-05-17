package algorithms

import "vengeful_replacer/app/dictionaries"

type Algorithm interface {
	Run()
	GetResult() string
	SetDictionary(dictionary dictionaries.Dictionary)
	SetData(string)
}
