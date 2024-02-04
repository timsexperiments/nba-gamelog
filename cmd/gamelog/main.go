package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/timsexperiments/nba-gamelog/internal/constants"
	"github.com/timsexperiments/nba-gamelog/internal/gamelog"
	"github.com/timsexperiments/nba-gamelog/internal/util"
)

func main() {
	season := flag.Int("season", time.Now().Year(), "The NBA season in YY or YYYY format (e.g., '23' or '2023')")
	startSeason := flag.Int("start", time.Now().Year(), "Start of the range of seasons")
	endSeason := flag.Int("end", time.Now().Year(), "End of the range of seasons")
	single := flag.Bool("single", false, "Whether to output a single file with the combined data for all seasons in the range.")
	dir := flag.String("dir", "nba_gamelog", "Output file location")

	flag.Parse()

	actualStart := util.MinInt(*startSeason, *season)
	actualEnd := util.MinInt(*endSeason, *season)

	util.PrintWarning("Warning: Due to Basketball Reference rate limits, only 20 requests can be made per minute.")
	totalSeasons := actualEnd - actualStart + 1
	totalTimeSeconds := totalSeasons * 30 * 60 / 20
	fmt.Printf("Estimated processing time: %d minutes and %d seconds.\n", totalTimeSeconds/60, totalTimeSeconds%60)

	done := make(chan bool)
	go util.LoadingAnimation(done, "Creating Gamelog")

	startTime := time.Now()
	gamelogs, err := gamelog.SeasonsGamelog(actualStart, actualEnd, *dir, !*single)
	timeToComplete := time.Since(startTime)
	done <- true
	if err != nil {
		log.Fatalf("There was an issue getting the gamelog from Basketball Reference: %s", err)
	}

	if *single {
		util.WriteCSV(gamelogs, fmt.Sprintf("%s/%s", *dir, constants.DefaultFileName))
	}
	fmt.Printf("\nSuccessfully created gamelog at %s in %v.\n", *dir, timeToComplete)
}
