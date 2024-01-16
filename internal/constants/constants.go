package constants

const (
	ATL string = "ATL"
	BOS        = "BOS"
	BKN        = "BRK"
	CHA        = "CHO"
	CHI        = "CHI"
	CLE        = "CLE"
	DAL        = "DAL"
	DEN        = "DEN"
	DET        = "DET"
	GSW        = "GSW"
	HOU        = "HOU"
	IND        = "IND"
	LAC        = "LAC"
	LAL        = "LAL"
	MEM        = "MEM"
	MIA        = "MIA"
	MIL        = "MIL"
	MIN        = "MIN"
	NOP        = "NOP"
	NYK        = "NYK"
	OKC        = "OKC"
	ORL        = "ORL"
	PHI        = "PHI"
	PHO        = "PHO"
	POR        = "POR"
	SAC        = "SAC"
	SAS        = "SAS"
	TOR        = "TOR"
	UTA        = "UTA"
	WAS        = "WAS"
)

var TEAMS = []string{
	ATL, BOS, BKN,
	CHA, CHI, CLE,
	DAL, DEN, DET,
	GSW, HOU, IND,
	LAC, LAL, MEM,
	MIA, MIL, MIN,
	NOP, NYK, OKC,
	ORL, PHI, PHO,
	POR, SAC, SAS,
	TOR, UTA, WAS,
}

func ToTeamDisplay(team string) string {
	switch team {
	case "CHO":
		return "CHA"
	case "NJN":
		return "BKN"
	case "SEA":
		return "OKC"
	case "NOH":
		return "NOP"
	case "VAN":
		return "MEM"
	default:
		return team
	}
}
