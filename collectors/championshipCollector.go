package collector

import (
	"log"

	"github.com/gocolly/colly"
)

type ChampionshipCollectorInterface interface {
	CollectorInterface
	GetTeamPages() [20]string
}

type ChampionshipCollector struct {
	Url       string
	TeamPages [20]string
}

func (c *ChampionshipCollector) Collect() {
	collector := CreateCollector(c.Url)
	collector.OnHTML("table", func(e *colly.HTMLElement) {
		index := 0
		e.ForEach("tbody tr", func(_ int, el *colly.HTMLElement) {
			if index < 20 {
				link := el.ChildAttr("td:nth-child(2) a", "href")
				c.TeamPages[index] = link
				index++
			}
		})
	})
	collector.Visit(c.Url)
}

func (c ChampionshipCollector) GetTeamPages() [20]string {
	return c.TeamPages
}

func CollectChampionship() [20]string {
	var championshipCollector ChampionshipCollectorInterface = &ChampionshipCollector{CHAMPIONSHIP_PAGE_URL, [20]string{}}
	championshipCollector.Collect()
	log.Println("CollectChampionship()")
	return championshipCollector.GetTeamPages()
}
