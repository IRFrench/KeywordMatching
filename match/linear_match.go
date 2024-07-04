package match

import (
	"log/slog"
	"slices"
	"strings"
)

type LinearMatchClient struct {
	keywords []string
}

// Only works for one wildcard
func (l *LinearMatchClient) Match(keyword string) []string {
	lowerKeyword := strings.ToLower(keyword)

	// Split apart the word into const and var
	splitKeyword := strings.Split(lowerKeyword, "")

	// Rebuild the possible combanations into a list.
	var wildcardPosition int

	for index := range splitKeyword {
		if slices.Contains(ALPHABET, splitKeyword[index]) {
			continue
		}
		wildcardPosition = index
		break
	}

	possibleKeywords := []string{}
	for letterIndex := range ALPHABET {
		tempKeyword := make([]string, len(splitKeyword))
		for index := range splitKeyword {
			if index == wildcardPosition {
				tempKeyword[index] = ALPHABET[letterIndex]
				continue
			}
			tempKeyword[index] = splitKeyword[index]
		}
		possibleKeywords = append(possibleKeywords, strings.Join(tempKeyword, ""))
	}

	// Check if any of the rebuilt words are in the list.
	foundWords := []string{}
	for keywordIndex := range possibleKeywords {
		if slices.Contains(l.keywords, possibleKeywords[keywordIndex]) {
			foundWords = append(foundWords, possibleKeywords[keywordIndex])
		}
	}

	return foundWords

}

func (l *LinearMatchClient) Stat() {
	slog.Info("linear client stats", "keyword list length", len(l.keywords))
}

func NewLinearMatchClient() (LinearMatchClient, error) {
	keywordList, err := getKeywordList()
	if err != nil {
		return LinearMatchClient{}, err
	}

	return LinearMatchClient{
		keywords: keywordList,
	}, nil
}
