package collector

import (
	"log"

	"github.com/gocolly/colly"
)

type CollectorInterface interface {
	Collect() // needs to implement OnHtml and then call on visit
}

// creates a collector from a collector struct
func CreateCollector(url string) *colly.Collector {
	log.Println("Creating collector...")
	c := colly.NewCollector(
	//colly.AllowedDomains("transfermarkt.com", "transfermarkt.com.br"),
	)

	log.Println("Collector collecting at ", url, "...")

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "1 Mozilla/5.0 (iPad; CPU OS 12_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148")
		log.Println("Visiting: ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("Page visited: ", r.Request.URL)
	})

	// c.OnHTML("a", func(e *colly.HTMLElement) {
	// 	// printing all URLs associated with the a links in the page
	// 	log.Println("%v", e.Attr("title"))
	// 	//log.Println("%v", e.Attr("href"))
	// })

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnScraped(func(r *colly.Response) {
		log.Println(r.Request.URL, " scraped!")
	})

	return c

	//c.Visit(collector.Url)
}
