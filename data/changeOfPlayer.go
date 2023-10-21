package data

type ChangeOfPlayer struct {
	ID             int64
	Minute         uint8
	PlayerLeaving  Player
	PlayerEntering Player
}
