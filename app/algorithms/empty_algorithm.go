package algorithms

import "vengeful_replacer/app/dictionaries"

type EmptyAlgorithm struct {
	data       string
	result     string
	dictionary dictionaries.Dictionary
}

func (alg *EmptyAlgorithm) Run() {
	alg.result = alg.data
}

func (alg EmptyAlgorithm) GetResult() string {
	return alg.result
}

func New(data string) EmptyAlgorithm {
	newEntity := EmptyAlgorithm{data: data}
	newEntity.Run()
	return newEntity
}
