package collector

import (
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
		index := 0
		e.ForEach("tbody tr", func(_ int, el *colly.HTMLElement) {
			if index < 20 {
				name := el.ChildAttr("td:nth-child(2) a", "title")
				link := el.ChildAttr("td:nth-child(2) a", "href")
				teamPages[index] = TeamPage{
					name, link,
				}
				index++
			}
		})
	})

	collector.Visit(c.Url)
}
func CollectChampionship() {
	var championshipCollector CollectorInterface = ChampionshipCollector{CHAMPIONSHIP_PAGE_URL}
	championshipCollector.Collect()
}
