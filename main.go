package main

import (
	"fmt"

	collector "github.com/tsmurer/brasileirao-data/collectors"
)

func main() {
	listOfTeamLinks := collector.CollectChampionship()
	for _, teamLink := range listOfTeamLinks {
		fmt.Printf("- %s\n", teamLink)
		team := collector.CollectClub(teamLink)
		fmt.Println(team)
	}
}
