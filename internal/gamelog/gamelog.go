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
	"team", "season", "g", "date",
	"home_away", "team_wl", "team_tm",
	"team_opp", "team_fg", "team_fga",
	"team_fgp", "team_3p", "team_3pa",
	"team_3pp", "team_ft", "team_fta",
	"team_ftp", "team_orb", "team_trb",
	"team_ast", "team_stl", "team_blk",
	"team_tov", "team_pf", "opp_fg",
	"opp_fga", "opp_fgp", "opp_3p",
	"opp_3pa", "opp_3pp", "opp_ft",
	"opp_fta", "opp_ftp", "opp_orb",
	"opp_trb", "opp_ast", "opp_stl",
	"opp_blk", "opp_tov", "opp_pf",
}

func SeasonsGamelog(start, end int) ([][]string, error) {
	allContents := make([][]string, 0)

	callsLeftInLimit := brLimitPerMinute
	for season := end; season >= start; season-- {
		for _, team := range constants.TEAMS {
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
			allContents = append(allContents, seasonGamelog...)
			if callsLeftInLimit == 0 {
				callsLeftInLimit = brLimitPerMinute
				time.Sleep(time.Minute)
			}
		}
	}

	return append([][]string{headers}, allContents...), nil
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
