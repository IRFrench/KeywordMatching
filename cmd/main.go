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
		hashClient, err := match.NewLinearMatchClient()
		if err != nil {
			slog.Error("failed to build client", "error", err)
		}
		if err := runMatchClient(keyword, &hashClient); err != nil {
			slog.Error("failed to run client", "error", err)
		}
		return
	}

	if checks.regex {
		slog.Info("Running regex check")
		hashClient, err := match.NewRegexClient()
		if err != nil {
			slog.Error("failed to build client", "error", err)
		}
		if err := runMatchClient(keyword, &hashClient); err != nil {
			slog.Error("failed to run client", "error", err)
		}
		return
	}

	if checks.rollingHash {
		slog.Info("Running rolling hash check")
		hashClient, err := match.NewRollingHashClient()
		if err != nil {
			slog.Error("failed to build client", "error", err)
		}
		if err := runMatchClient(keyword, &hashClient); err != nil {
			slog.Error("failed to run client", "error", err)
		}
		return
	}

	if checks.altRollingHash {
		slog.Info("Running alternate rolling hash check")
		hashClient, err := match.NewAltRollingHashClient()
		if err != nil {
			slog.Error("failed to build client", "error", err)
		}
		if err := runMatchClient(keyword, &hashClient); err != nil {
			slog.Error("failed to run client", "error", err)
		}
		return
	}

	if checks.runeRollingHash {
		slog.Info("Running rune rolling hash check")
		hashClient, err := match.NewRuneRollingHashClient()
		if err != nil {
			slog.Error("failed to build client", "error", err)
		}
		if err := runMatchClient(keyword, &hashClient); err != nil {
			slog.Error("failed to run client", "error", err)
		}
		return
	}

	slog.Warn("No checks requested")
}

type checks struct {
	linear          bool
	regex           bool
	rollingHash     bool
	altRollingHash  bool
	runeRollingHash bool
}

func getFlags() (string, checks) {
	keyword := flag.String("k", "", "keyword")
	linearFlag := flag.Bool("l", false, "run linear matching")
	regexFlag := flag.Bool("r", false, "run regex matching")
	rollingHashFlag := flag.Bool("rh", false, "run rolling hash matching")
	altRollingHashFlag := flag.Bool("arh", false, "run alt rolling hash matching")
	runeRollingHashFlag := flag.Bool("rrh", false, "run rune rolling hash matching")
	flag.Parse()

	return *keyword, checks{
		linear:          *linearFlag,
		regex:           *regexFlag,
		rollingHash:     *rollingHashFlag,
		altRollingHash:  *altRollingHashFlag,
		runeRollingHash: *runeRollingHashFlag,
	}
}

func runMatchClient(keyword string, client match.KeywordMatchClient) error {
	client.Stat()
	startTime := time.Now()
	matches, err := client.Match(keyword)
	if err != nil {
		return err
	}
	endTime := time.Now()
	slog.Info("found matches", "time taken", endTime.Sub(startTime), "matches", len(matches))
	slog.Info("results", "matches", matches)
	return nil
}
