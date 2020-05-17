package dictionaries

type SimpleDictionary struct {
	maintenance map[string]string
}

func (dic SimpleDictionary) GetMap() map[string]string {
	return dic.maintenance
}

func NewSimpleDictionary() *SimpleDictionary {

	easyDict := map[string]string{
		"а": "a", // cyrillic:latin
		"с": "c",
		"е": "e",
		"о": "o",
		"р": "p",
		"к": "k",
		"х": "x",
		"у": "y",
	}

	newEntity := SimpleDictionary{maintenance: easyDict}

	return &newEntity
}
