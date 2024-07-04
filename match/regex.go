package match

import (
	"fmt"
	"log/slog"
	"regexp"
	"slices"
	"strings"
)

type RegexClient struct {
	keywords []string
}

func (r *RegexClient) Match(keyword string) ([]string, error) {
	lowerKeyword := strings.ToLower(keyword)

	// Turn wildcard to wildcard
	splitKeyword := strings.Split(lowerKeyword, "")
	for index := range splitKeyword {
		if !slices.Contains(ALPHABET, splitKeyword[index]) {
			splitKeyword[index] = "."
		}
	}

	// Create regex pattern
	keywordPattern := fmt.Sprintf("^%v$", strings.Join(splitKeyword, ""))

	// Check all of the keywords
	matchedKeywords := []string{}
	for keywordIndex := range r.keywords {
		match, err := regexp.MatchString(keywordPattern, r.keywords[keywordIndex])
		if err != nil {
			return nil, err
		}
		if match {
			matchedKeywords = append(matchedKeywords, r.keywords[keywordIndex])
		}
	}

	return matchedKeywords, nil
}

func (r *RegexClient) Stat() {
	slog.Info("linear client stats", "keyword list length", len(r.keywords))
}

func NewRegexClient() (RegexClient, error) {
	keywordList, err := getKeywordList()
	if err != nil {
		return RegexClient{}, err
	}

	return RegexClient{
		keywords: keywordList,
	}, nil
}
