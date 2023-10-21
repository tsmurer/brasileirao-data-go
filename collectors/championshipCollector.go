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
	// teamPages := [20]TeamPage{}
	collector.OnHTML("td.hauptlink", func(e *colly.HTMLElement) {
		// printing all URLs associated with the a links in the page
		// if strings.Contains(e.ChildAttr("a", "href"), "verein") {
		// 	log.Println("value found: %v", e.ChildAttr("a", "title"), e.ChildAttr("a", "href"))
		// 	teamPage := TeamPage{e.ChildAttr("a", "title"), e.ChildAttr("a", "href")}
		// 	for i:=0; i < 20; i++{
		// 		if(len)
		// 	}
		// }

	})

	collector.Visit(c.Url)
}
func CollectChampionship() {
	var championshipCollector CollectorInterface = ChampionshipCollector{CHAMPIONSHIP_PAGE_URL}
	championshipCollector.Collect()
}
