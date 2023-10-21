package data

type Player struct {
	ID            int64
	Name          string
	Team          Team
	goals         uint16
	assists       uint16
	matchesPlayed uint16
	yellowCards   uint16
	redCards      uint16
}
