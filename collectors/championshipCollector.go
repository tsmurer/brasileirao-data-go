package collector

func CollectChampionship() {
	var championshipCollector = Collector{CHAMPIONSHIP_URL}
	championshipCollector.Collect()
}
