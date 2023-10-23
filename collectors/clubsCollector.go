package collector

type ClubCollector struct {
	Url string
}

type ClubPage struct {
	Name string
	Url  string
}

func (c ClubCollector) Collect() {
	collector := CreateCollector(c.Url)

	collector.Visit(c.Url)
}
func CollectClub(clubUrl string) {
	var clubCollector CollectorInterface = ClubCollector{clubUrl}
	clubCollector.Collect()
}
