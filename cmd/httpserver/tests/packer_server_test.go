package tests_test

import (
	"fmt"
	"github.com/vcayetano/fulfillment/acceptancetests"
	"github.com/vcayetano/fulfillment/acceptancetests/httpserver"
	"github.com/vcayetano/fulfillment/specifications"
	"net/http"
	"testing"
	"time"
)

func TestPackerServer(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	const PORT = "8080"
	acceptancetests.StartDockerTestServer(t, PORT)
	client := &http.Client{Timeout: 3 * time.Second}
	driver := httpserver.Driver{BaseURL: fmt.Sprintf("http://localhost:%s", PORT), Client: client}
	specifications.PackerSpecification(t, driver)
}
