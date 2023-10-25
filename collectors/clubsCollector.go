package collector

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type ClubCollectorInterface interface {
	CollectorInterface
	GetTeam() Team
}

type ClubCollector struct {
	Team
	Url string
}

type ClubPage struct {
	Name string
	Url  string
}

type Team struct {
	Name         string
	ScoredPoints string
	GoalsAgainst string
	GoalsFor     string
	Matches      []Match
}

type Match struct {
	AwayTeamName  string
	HomeTeamName  string
	AwayTeamScore string
	HomeTeamScore string
}

func (c ClubCollector) Collect() {
	collector := CreateCollector(c.Url)
	// wanted data:
	// name
	// matches played
	//startseite -> spielplan
	fmt.Println(c.Url)
	team := Team{}
	collector.OnHTML("header.data-header", func(e *colly.HTMLElement) {
		team.Name = strings.TrimSpace(e.ChildText("h1.data-header__headline-wrapper"))
		team.ScoredPoints = strings.TrimSpace(e.ChildText("a.data-header__market-value-wrapper"))
	})

	// Set up the callback to extract match information
	collector.OnHTML("div.responsive-table tbody tr", func(e *colly.HTMLElement) {
		match := Match{
			AwayTeamName:  strings.TrimSpace(e.ChildText("td.no-border-links.hauptlink a")),
			HomeTeamName:  strings.TrimSpace(e.ChildText("td.no-border-rechts a")),
			AwayTeamScore: strings.TrimSpace(e.ChildText("td.zentriert.hauptlink")),
			HomeTeamScore: strings.TrimSpace(e.ChildText("td.zentriert.no-border-links")),
		}
		team.Matches = append(team.Matches, match)
	})

	// Navigate to the URL you want to scrape
	err := collector.Visit(c.Url)

	if err != nil {
		fmt.Println("Error:", err)
	}

	// Print the extracted team and match data
	fmt.Printf("Team Name: %s\n", team.Name)
	fmt.Printf("Scored Points: %s\n", team.ScoredPoints)

	fmt.Println("Matches:")
	for i, match := range team.Matches {
		fmt.Printf("Match %d\n", i+1)
		fmt.Printf("Away Team: %s\n", match.AwayTeamName)
		fmt.Printf("Home Team: %s\n", match.HomeTeamName)
		fmt.Printf("Away Team Score: %s\n", match.AwayTeamScore)
		fmt.Printf("Home Team Score: %s\n", match.HomeTeamScore)
	}
}

func (c ClubCollector) GetTeam() Team {
	return c.Team
}

func CollectClub(clubUrl string) Team {
	url := "https://www.transfermarkt.com" + strings.Replace(clubUrl, "startseite", "spielplan", 1)
	var clubCollector ClubCollectorInterface = &ClubCollector{Team{}, url}
	clubCollector.Collect()
	return clubCollector.GetTeam()
}
