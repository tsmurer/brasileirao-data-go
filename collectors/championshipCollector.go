package collector

import (
	"log"

	"github.com/gocolly/colly"
)

type ChampionshipCollector struct {
	Url string
}

type TeamPage struct {
	Name string
	Url  string
}

func (c ChampionshipCollector) Collect() {
	collector := CreateCollector(c.Url)
	teamPages := [20]TeamPage{}
	collector.OnHTML("table", func(e *colly.HTMLElement) {
		// Initialize an index
		index := 0

		// Find and iterate through table rows
		e.ForEach("tbody tr", func(_ int, el *colly.HTMLElement) {
			// Skip the first row (table header)
			if index < 20 {
				// Extract club name and link from the table data
				name := el.ChildAttr("td:nth-child(2) a", "title")
				link := el.ChildAttr("td:nth-child(2) a", "href")

				// Add the club to the fixed-size array
				teamPages[index] = TeamPage{
					name, link,
				}

				// Increment the index
				index++
			}
		})
	})

	collector.Visit(c.Url)
	for i := 0; i < 20; i++ {
		log.Println("ARRAY:")
		log.Println(i, ": ", teamPages[i])
	}
}
func CollectChampionship() {
	var championshipCollector CollectorInterface = ChampionshipCollector{CHAMPIONSHIP_PAGE_URL}
	championshipCollector.Collect()
}
