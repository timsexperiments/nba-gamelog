package gamelog

import (
	"fmt"
	"strings"
	"time"

	"github.com/timsexperiments/nba-gamelog/internal/constants"
	"github.com/timsexperiments/nba-gamelog/internal/scraper"
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

func SeasonsGamelog(start, end int) ([][]string, error) {
	allContents := make([][]string, 0)

	callsLeftInLimit := brLimitPerMinute
	for season := end; season >= start; season-- {
		for i, team := range constants.TEAMS {
			callsLeftInLimit--
			if team == constants.BKN && season < 2013 {
				team = "NJN"
			}
			if team == constants.OKC && season < 2009 {
				team = "SEA"
			}
			if team == constants.NOP && season < 2013 {
				team = "NOH"
			}
			if team == constants.MEM && season < 2002 {
				team = "VAN"
			}
			if team == constants.CHA && season < 2013 {
				team = "CHA"
			}
			seasonGamelog, err := seasonLog(team, season)
			if err != nil {
				return nil, err
			}
			allContents = append(allContents, transformSeasonGamelog(seasonGamelog)...)
			if callsLeftInLimit == 0 && (season != end || i != len(constants.TEAMS)-1) {
				callsLeftInLimit = brLimitPerMinute
				time.Sleep(time.Minute)
			}
		}
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
