package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/timsexperiments/nba-gamelog/internal/gamelog"
	"github.com/timsexperiments/nba-gamelog/internal/util"
)

func main() {
	season := flag.Int("season", 2024, "The NBA season in YY or YYYY format (e.g., '23' or '2023')")
	startSeason := flag.Int("start", 2024, "Start of the range of seasons")
	endSeason := flag.Int("end", 0, "End of the range of seasons")
	output := flag.String("output", "", "Output file location")

	flag.Parse()

	// Default end season to the current year if not provided
	if *endSeason == 0 {
		*endSeason = time.Now().Year()
	}

	if *output == "" {
		homeDir, _ := os.UserHomeDir()
		defaultOutput := filepath.Join(homeDir, "nba_gamelog.csv")
		output = &defaultOutput
	}

	util.PrintWarning("Warning: Due to Basketball Reference rate limits, only 20 requests can be made per second.")
	actualStart := util.MinInt(*startSeason, *season)
	totalSeasons := *endSeason - actualStart + 1
	totalTimeSeconds := totalSeasons * 30 * 60 / 20
	fmt.Printf("Estimated processing time: %d minutes and %d seconds.\n", totalTimeSeconds/60, totalTimeSeconds%60)

	gamelogs, err := gamelog.SeasonsGamelog(actualStart, *endSeason)
	if err != nil {
		log.Fatalf("There was an issue getting the gamelog from Basketball Reference: %s", err)
	}
	util.WriteCSV(gamelogs, *output)
	fmt.Printf("\nSuccessfully created gamelog at %s.\n", *output)
}
