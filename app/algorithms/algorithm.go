package algorithms

type Algorithm interface {
	Run()
	GetResult() string
}
