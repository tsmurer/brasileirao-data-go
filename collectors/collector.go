package collector

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Collector struct {
	Url string
}

// type CollectorInterface interface {
// 	(*Collector)Collect(string)
// }

func (collector *Collector) Collect() {
	log.Println("Collecting...")
	c := colly.NewCollector()
	c.Visit(collector.Url)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})
}
