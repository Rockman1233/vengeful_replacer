package dictionaries

type Dictionary interface {
	getMap() map[string]string
}
