package main

import (
	"flag"
	"keyword-matching/match"
	"log/slog"
	"time"
)

func main() {
	keyword, checks := getFlags()

	if keyword == "" {
		slog.Warn("missing keyword (-k)")
	}

	if checks.linear {
		slog.Info("Running linear check")
		linearClient, err := match.NewLinearMatchClient()
		if err != nil {
			slog.Error("failed to build client", "error", err)
		}
		linearClient.Stat()
		startTime := time.Now()
		matches := linearClient.Match(keyword)
		endTime := time.Now()
		slog.Info("found matches", "time taken", endTime.Sub(startTime), "matches", len(matches))
		return
	}

	if checks.regex {
		slog.Info("Running regex check")
		regexClient, err := match.NewRegexClient()
		if err != nil {
			slog.Error("failed to build client", "error", err)
		}
		regexClient.Stat()
		startTime := time.Now()
		matches, err := regexClient.Match(keyword)
		endTime := time.Now()
		if err != nil {
			slog.Error("error finding matches", "error", err)
		}
		slog.Info("found matches", "time taken", endTime.Sub(startTime), "matches", len(matches))
		return
	}

	if checks.rollingHash {
		slog.Info("Running rolling hash check")
		hashClient, err := match.NewRollingHashClient()
		if err != nil {
			slog.Error("failed to build client", "error", err)
		}
		hashClient.Stat()
		startTime := time.Now()
		matches := hashClient.Match(keyword)
		endTime := time.Now()
		slog.Info("found matches", "time taken", endTime.Sub(startTime), "matches", len(matches))
		return
	}

	slog.Warn("No checks requested")
	return
}

type checks struct {
	linear      bool
	regex       bool
	rollingHash bool
}

func getFlags() (string, checks) {
	keyword := flag.String("k", "", "keyword")
	linearFlag := flag.Bool("l", false, "run linear matching")
	regexFlag := flag.Bool("r", false, "run regex matching")
	rollingHashFlag := flag.Bool("rh", false, "run rolling hash matching")
	flag.Parse()

	return *keyword, checks{
		linear:      *linearFlag,
		regex:       *regexFlag,
		rollingHash: *rollingHashFlag,
	}
}
