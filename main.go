package main

import (
	"fmt"

	collector "github.com/tsmurer/brasileirao-data/collectors"
)

func main() {
	listOfTeamLinks := collector.CollectChampionship()
	teams := []collector.Team{}
	for _, teamLink := range listOfTeamLinks {
		team := collector.CollectClub(teamLink)
		teams = append(teams, team)
	}
	fmt.Println(teams)
}
