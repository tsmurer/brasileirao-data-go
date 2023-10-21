package data

type Team struct {
	ID            int64
	Name          string
	Players       []Player
	MatchesPlayed uint16
	Score         uint16
	GoalsFor      uint16
	GoalsAgainst  uint16
	YellowCards   uint16
	RedCards      uint16
}
