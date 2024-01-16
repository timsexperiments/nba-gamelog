package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/timsexperiments/nba-gamelog/internal/scraper"
)

func main() {
	season := flag.String("season", "", "The NBA season in YY or YYYY format (e.g., '23' or '2023')")
	startSeason := flag.String("start", "", "Start of the range of seasons")
	endSeason := flag.String("end", "", "End of the range of seasons")
	output := flag.String("output", "", "Output file location")

	flag.Parse()

	// Default end season to the current year if not provided
	currentYear := time.Now().Year()
	if *endSeason == "" {
		*endSeason = fmt.Sprintf("%d", currentYear)
	}

	if *season == "" && *startSeason == "" {
		log.Fatalf("Error: You must specify either a single season or a start season.")
	}

	if *output == "" {
		homeDir, _ := os.UserHomeDir()
		defaultOutput := filepath.Join(homeDir, "nba_gamelog.csv")
		output = &defaultOutput
	}

	fmt.Printf("Season: %s, Start: %s, End: %s, Output: %s\n", *season, *startSeason, *endSeason, *output)

	contents, err := scraper.LoadTeamSeasonLog("ATL", 2024)
	if err != nil {
		log.Fatalf("There was an issue getting the gamelog from Basketball Reference: %s", err)
	}
	fmt.Println(contents)
}
