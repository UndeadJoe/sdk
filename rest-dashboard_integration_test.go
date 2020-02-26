package sdk_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/grafana-tools/sdk"
)

func Test_Dashboard_CRUD(t *testing.T) {
	var (
		boardLinks []sdk.FoundBoard
	)

	shouldSkip(t)

	client := getClient()

	var board sdk.Board
	raw, _ := ioutil.ReadFile("testdata/new-empty-dashboard-2.6.json")

	err := json.Unmarshal(raw, &board)

	if _, err := client.SetDashboard(board, true); err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error(err)
	}

	if boardLinks, err = client.SearchDashboards("", false); err != nil {
		t.Error(err)
	}

	for _, link := range boardLinks {
		_, _, err := client.GetDashboardByUID(link.UID)
		if err != nil {
			t.Error(err)
		}
	}
}
