package scanner

type Scanner interface {
	GetData() string
	TakeData() string
}
