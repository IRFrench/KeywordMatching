package bench

import (
	"keyword-matching/match"
	"testing"
)

func BenchmarkRegexMatch(b *testing.B) {
	regexClient, err := match.NewRegexClient()
	if err != nil {
		panic(err)
	}

	testCases := map[string]string{
		"keyword test":  "p*rn",
		"short keyword": "c*t",
		"long keyword":  "av*cado",
	}

	for name, keyword := range testCases {
		b.Run(name, func(b *testing.B) {
			regexClient.Match(keyword)
		})
	}
}
