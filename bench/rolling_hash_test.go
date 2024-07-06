package bench

import (
	"keyword-matching/match"
	"testing"
)

func BenchmarkRollingHashMatch(b *testing.B) {
	rollingHashClient, err := match.NewRollingHashClient()
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
			rollingHashClient.Match(keyword)
		})
	}
}

func BenchmarkAltRollingHashMatch(b *testing.B) {
	rollingHashClient, err := match.NewAltRollingHashClient()
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
			rollingHashClient.Match(keyword)
		})
	}
}

func BenchmarkRuneRollingHashMatch(b *testing.B) {
	rollingHashClient, err := match.NewRuneRollingHashClient()
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
			rollingHashClient.Match(keyword)
		})
	}
}
