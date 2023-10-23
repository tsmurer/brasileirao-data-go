package collector

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tsmurer/brasileirao-data/assets"
)

func setupSuite() {

}

func TestChampionshipCollector_Collect(t *testing.T) {
	embeddedHTML := assets.Assets

	htmlData, err := embeddedHTML.ReadFile("championshipPage.html")
	if err != nil {
		t.Fatalf("Error reading embedded file: %v", err)
	}

	// Create a test HTTP server to serve the HTML content
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(htmlData))
	}))
	defer server.Close()

	collector := &ChampionshipCollector{Url: server.URL, TeamPages: [20]TeamPage{}}
	collector.Collect()

	for _, page := range collector.TeamPages {
		if page.Name == "" || page.Url == "" {
			t.Errorf("Expected all indexes populated but found: { Name: %v, Url: %v }", page.Name, page.Url)
		}
	}
}

// Add more test functions for other aspects of your code, if necessary

func TestChampionshipCollector_GetTeamPages(t *testing.T) {
	collector := ChampionshipCollector{
		TeamPages: [20]TeamPage{ /* Initialize with your test data */ },
	}
	result := collector.GetTeamPages()
	if len(result) != 20 {
		t.Errorf("Expected result length of 20, but got %d", len(result))
	}
}
