package data

type Championship struct {
	ID     int64
	Teams  [20]Team
	Rounds [38]Round
}
