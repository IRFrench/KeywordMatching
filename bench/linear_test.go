package bench

import (
	"keyword-matching/match"
	"testing"
)

func BenchmarkLinearMatch(b *testing.B) {
	linearClient, err := match.NewLinearMatchClient()
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
			linearClient.Match(keyword)
		})
	}
}
