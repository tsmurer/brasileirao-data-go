package collector

import (
	"log"

	"github.com/gocolly/colly"
)

type ChampionshipCollector struct {
	Url string
}

func (c ChampionshipCollector) Collect() {
	collector := CreateCollector(c.Url)
	collector.OnHTML("a", func(e *colly.HTMLElement) {
		// printing all URLs associated with the a links in the page
		log.Println("%v", e.Attr("href"))
	})

	collector.Visit(c.Url)
}
func CollectChampionship() {
	var championshipCollector CollectorInterface = ChampionshipCollector{CHAMPIONSHIP_PAGE_URL}
	championshipCollector.Collect()

}
