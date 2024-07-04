package match

import (
	"os"
	"strings"
)

var (
	ALPHABET = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "'", "ó", "ñ", "é", "ç", "å", "ö", "ê", "è", "ä", "û", "â", "ü", "á", "í", "ô"}
	ATOI     = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26, "'": 27, "ó": 28, "ñ": 29, "é": 30, "ç": 31, "å": 32, "ö": 33, "ê": 34, "è": 35, "ä": 36, "û": 37, "â": 38, "ü": 39, "á": 40, "í": 41, "ô": 42}
)

func getKeywordList() ([]string, error) {
	keywords, err := os.ReadFile("/usr/share/dict/words")
	if err != nil {
		return nil, err
	}

	keywordList := []string{}
	for _, keyword := range strings.Split(string(keywords), "\n") {
		keywordList = append(keywordList, strings.ToLower(keyword))
	}

	return keywordList, nil
}
