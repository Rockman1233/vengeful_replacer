package dictionaries

type SimpleDictionary struct {
	maintenance map[string]string
}

func (dic SimpleDictionary) getMap() map[string]string {
	return dic.maintenance
}

func New() (SimpleDictionary, err) {

	easyDict := map[string]string{
		"а": "а", // cyrillic:latin
		"с": "c",
		"е": "e",
		"о": "o",
		"р": "p",
		"к": "k",
		"х": "x",
		"у": "y",
	}

	newEntity := SimpleDictionary{maintenance: easyDict}

	return newEntity, err
}
