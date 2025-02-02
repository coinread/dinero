package dinero

import (
	"os"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

// TestListCurrencies will test listing currencies from the OXR api.
func TestListCurrencies(t *testing.T) {
	// Register the test.
	RegisterTestingT(t)

	// Init dinero client.
	client := NewClient(os.Getenv("OPEN_EXCHANGE_APP_ID"), "AUD", 1*time.Minute)

	// Get the currencies
	rsp, err := client.Currencies.List()
	if err != nil {
		t.Fatalf("Unexpected error running client.Currencies.Get(): %s", err.Error())
	}

	Expect(err).Should(BeNil())
	Expect(rsp).Should(ContainElement(&CurrencyResponse{
		Code: "AUD",
		Name: "Australian Dollar",
	}))
	Expect(rsp).Should(ContainElement(&CurrencyResponse{
		Code: "NZD",
		Name: "New Zealand Dollar",
	}))
}
