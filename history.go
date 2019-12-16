package dinero

import (
	"fmt"
	"strings"
	"time"
)

const (
	historyAPIPath = "historical/%s.json"
)

// CurrenciesService handles currency request/responses.
type HistoricalService struct {
	client *Client
}

// NewCurrenciesService creates a new handler for this service.
func NewHistoricalService(
	client *Client,
) *HistoricalService {
	return &HistoricalService{
		client,
	}
}

// HistoricalResponse represents the OXR response
// See
type HistoricalResponse struct {
	Disclaimer string             `json:"disclaimer"`
	License    string             `json:"license"`
	TimeStamp  int64              `json:"timestamp"`
	Base       string             `json:"base"`
	Rates      map[string]float64 `json:"rates"`
}

// Get will fetch all list of all currencies available via the OXR api.
// See https://docs.openexchangerates.org/docs/time-series-json
func (s *HistoricalService) Get(timeStamp time.Time, baseCurrency string, symbols []string, includeAlternativeRates bool) (*HistoricalResponse, error) {
	queryParams := make(map[string]string, 0)
	hourStamp := timeStamp.Truncate(time.Hour * 24).Format("2006-01-02")

	if baseCurrency != "" {
		queryParams["base"] = baseCurrency
	}

	if len(symbols) > 0 {
		queryParams["symbols"] = strings.Join(symbols, ",")
	}

	queryParams["show_alternative"] = fmt.Sprintf("%t", includeAlternativeRates)

	apiPath := fmt.Sprintf(historyAPIPath, hourStamp)
	for key, value := range queryParams {
		// Is there one in here already?
		if strings.Contains(apiPath, "?") {
			apiPath += "&"
		} else {
			apiPath += "?"
		}

		apiPath += fmt.Sprintf("%s=%s", key, value)
	}

	// Build request.
	req, err := s.client.NewRequest(
		"GET",
		apiPath,
		nil,
	)
	if err != nil {
		return nil, err
	}

	// Make request.
	rsp := new(HistoricalResponse)

	// Because rsp is already a pointer
	if _, err = s.client.Do(req, rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}
