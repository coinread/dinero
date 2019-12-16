package dinero

import (
	"os"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

// TestListCurrencies will test listing currencies from the OXR api.
func TestGetTimeSeries(t *testing.T) {
	// Register the test.
	RegisterTestingT(t)

	// Init dinero client.
	client := NewClient(os.Getenv("OPEN_EXCHANGE_APP_ID"), "AUD", 1*time.Minute)

	// Get the currencies
	rsp, err := client.TimeSeries.Get(time.Now().Add(time.Hour * -24), time.Now(), "", []string{}, false)
	if err != nil {
		t.Fatalf("Unexpected error running client.TimeSeries.Get(): %s", err.Error())
	}

	Expect(err).Should(BeNil())
	Expect(rsp).Should(Not(BeNil()))
	Expect(rsp.Base).ShouldNot(BeEmpty())
}
