package match

import (
	"log/slog"
	"slices"
	"strings"
)

type RollingHashClient struct {
	keywordMap   map[int][]string
	keywordCount int
}

// will only work for one wildcard
func (r *RollingHashClient) Match(keyword string) ([]string, error) {
	lowerKeyword := strings.ToLower(keyword)

	// Split apart the word into const and var
	splitKeyword := strings.Split(lowerKeyword, "")

	// Rebuild the possible combanations into a list.
	var wildcardPosition int

	for index := range splitKeyword {
		if slices.Contains(MATCH_CHARACTERS, splitKeyword[index]) {
			wildcardPosition = index
			break
		}
	}

	possibleKeywords := map[int]string{}
	for letterIndex := range ALPHABET {
		tempKeyword := make([]string, len(splitKeyword))
		for index := range splitKeyword {
			if index == wildcardPosition {
				tempKeyword[index] = ALPHABET[letterIndex]
				continue
			}
			tempKeyword[index] = splitKeyword[index]
		}
		combinedKeyword := strings.Join(tempKeyword, "")
		possibleKeywords[numberHash(combinedKeyword)] = combinedKeyword
	}

	// Check if any of the rebuilt words are in the list.
	foundWords := []string{}
	for hash, keyword := range possibleKeywords {
		shortList, found := r.keywordMap[hash]
		if !found {
			continue
		}
		if slices.Contains(shortList, keyword) {
			foundWords = append(foundWords, keyword)
		}
	}

	return foundWords, nil
}

func (r *RollingHashClient) Stat() {
	maxSize := 0
	minSize := 0
	totalSegments := 0
	medianMap := map[int]int{}

	for _, list := range r.keywordMap {
		listLength := len(list)

		if minSize == 0 {
			minSize = listLength
		}
		if listLength > maxSize {
			maxSize = listLength
		}
		if listLength < minSize {
			minSize = listLength
		}
		totalSegments++

		_, found := medianMap[listLength]
		if !found {
			medianMap[listLength] = 0
		}
		medianMap[listLength]++
	}

	highestCount := 0
	for _, count := range medianMap {
		if count > highestCount {
			highestCount = count
		}
	}

	allMedians := []int{}
	for length, count := range medianMap {
		if count == highestCount {
			allMedians = append(allMedians, length)
		}
	}

	medianMiddle := len(allMedians) / 2
	slices.Sort(allMedians)

	slog.Info("rolling hash client stats",
		"max segment size", maxSize,
		"min segment size", minSize,
		"avg segment size", r.keywordCount/totalSegments,
		"median segment size", allMedians[medianMiddle],
		"median count", highestCount,
		"all keywords", r.keywordCount,
		"total segments", totalSegments,
	)
}

func NewRollingHashClient() (RollingHashClient, error) {
	keywordList, err := getKeywordList()
	if err != nil {
		return RollingHashClient{}, err
	}

	keywordHashMap := map[int][]string{}
	for _, keyword := range keywordList {
		hashIndex := numberHash(keyword)
		_, found := keywordHashMap[hashIndex]
		if !found {
			keywordHashMap[hashIndex] = []string{keyword}
		}
		keywordHashMap[hashIndex] = append(keywordHashMap[hashIndex], keyword)
	}

	return RollingHashClient{
		keywordMap:   keywordHashMap,
		keywordCount: len(keywordList),
	}, nil
}

func numberHash(keyword string) int {
	splitKeyword := strings.Split(keyword, "")
	numberHash := 1
	for index := range splitKeyword {
		number, found := ATOI[splitKeyword[index]]
		if !found {
			slog.Warn("non known letter found", "word", keyword, "letter", splitKeyword[index])
		}
		numberHash *= number
	}
	return numberHash
}
