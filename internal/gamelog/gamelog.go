package gamelog

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/timsexperiments/nba-gamelog/internal/constants"
	"github.com/timsexperiments/nba-gamelog/internal/scraper"
	"github.com/timsexperiments/nba-gamelog/internal/util"
)

const brLimitPerMinute = 20

var headers = []string{
	"home", "season", "game",
	"date", "away", "team_wl",
	"team_score", "opp_score", "team_fg",
	"team_fga", "team_fgp", "team_3p",
	"team_3pa", "team_3pp", "team_ft",
	"team_fta", "team_ftp", "team_orb",
	"team_trb", "team_ast", "team_stl",
	"team_blk", "team_tov", "team_pf",
	"opp_fg", "opp_fga", "opp_fgp",
	"opp_3p", "opp_3pa", "opp_3pp",
	"opp_ft", "opp_fta", "opp_ftp",
	"opp_orb", "opp_trb", "opp_ast",
	"opp_stl", "opp_blk", "opp_tov",
	"opp_pf",
}

func SeasonsGamelog(start, end int, folder string, writeAll bool) ([][]string, error) {
	allContents := make([][]string, 0)

	callsLeftInLimit := brLimitPerMinute
	for season := end; season >= start; season-- {
		currentSeasonLog := make([][]string, 0)
		for i, team := range constants.TEAMS {
			callsLeftInLimit--
			ogTeam := team
			if ogTeam == constants.BKN && season < 2013 {
				team = "NJN"
			}
			if ogTeam == constants.OKC && season < 2009 {
				team = "SEA"
			}
			if ogTeam == constants.NOP && season < 2003 {
				continue
			}
			if ogTeam == constants.NOP && (season == 2006 || season == 2007) {
				team = "NOK"
			}
			if ogTeam == constants.NOP && season < 2014 {
				team = "NOH"
			}
			if ogTeam == constants.NOP && season < 2003 {
				team = "CHH"
			}
			if ogTeam == constants.MEM && season < 2002 {
				team = "VAN"
			}
			if ogTeam == constants.CHA && season < 2005 {
				continue
			}
			if ogTeam == constants.CHA && season < 2015 {
				team = "CHA"
			}

			seasonGamelog, err := seasonLog(team, season)
			for try := 0; try < 2; try++ {
				if err != nil {
					seasonGamelog, err = seasonLog(team, season)
					callsLeftInLimit--
				}
			}
			if err != nil {
				return nil, err
			}

			currentSeasonLog = append(currentSeasonLog, transformSeasonGamelog(seasonGamelog)...)
			if callsLeftInLimit == 0 && (season != end || i != len(constants.TEAMS)-1) {
				callsLeftInLimit = brLimitPerMinute
				time.Sleep(time.Minute)
			}
		}
		if writeAll {
			writeFile := fmt.Sprintf("%s/%d_%s", folder, season, constants.DefaultFileName)
			if err := util.WriteCSV(append([][]string{headers}, currentSeasonLog...), writeFile); err != nil {
				fmt.Fprintf(os.Stderr, "Unable to write file %s.\n", err)
			}
		}

		allContents = append(allContents, currentSeasonLog...)
	}

	return append([][]string{headers}, allContents...), nil
}

func transformSeasonGamelog(gamelog [][]string) [][]string {
	tansformedGamelog := gamelog[6:]
	for i, gameData := range tansformedGamelog {
		tansformedGamelog[i] = append(append(gameData[0:4], gameData[5:25]...), gameData[26:]...)
	}
	return tansformedGamelog
}

func seasonLog(team string, season int) ([][]string, error) {
	transformedContents := make([][]string, 0)
	contents, err := scraper.LoadTeamSeasonLog(team, season)
	if err != nil {
		return nil, err
	}

	for _, row := range strings.Split(contents, "\n") {
		rowContents := strings.Split(row, ",")
		transformedContents = append(transformedContents, append(
			[]string{constants.ToTeamDisplay(team), fmt.Sprintf("%d", season)}, rowContents[1:]...))
	}

	return transformedContents, nil
}
