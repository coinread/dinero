package dinero

import (
	"fmt"
	"strings"
	"time"
)

const (
	timeSeriesAPIPath = "time-series.json"
)

// CurrenciesService handles currency request/responses.
type TimeSeriesService struct {
	client *Client
}

// NewCurrenciesService creates a new handler for this service.
func NewTimeSeriesService(
	client *Client,
) *TimeSeriesService {
	return &TimeSeriesService{
		client,
	}
}

// TimeSeriesResponse represents the OXR response
// See
type TimeSeriesResponse struct {
	Disclaimer string                        `json:"disclaimer"`
	License    string                        `json:"license"`
	StartDate  string                        `json:"start_date"`
	EndDate    string                        `json:"end_date"`
	Base       string                        `json:"base"`
	Rates      map[string]map[string]float64 `json:"rates"`
}

// Get will fetch all list of all currencies available via the OXR api.
// See https://docs.openexchangerates.org/docs/time-series-json
func (s *TimeSeriesService) Get(startTime, endTime time.Time, baseCurrency string, symbols []string, includeAlternativeRates bool) (*TimeSeriesResponse, error) {
	queryParams := make(map[string]string, 0)
	timeDeltaHours := endTime.Sub(startTime).Hours()
	timeDeltaDays := timeDeltaHours / 24
	if timeDeltaDays > 31 {
		return nil, fmt.Errorf("longest time delta supported is one month (31 days), you asked for %f day(s)", timeDeltaDays)
	}

	queryParams["start"] = startTime.Truncate(time.Hour * 24).Format("2006-01-02")
	queryParams["end"] = endTime.Truncate(time.Hour * 24).Format("2006-01-02")
	if baseCurrency != "" {
		queryParams["base"] = baseCurrency
	}

	if len(symbols) > 0 {
		queryParams["symbols"] = strings.Join(symbols, ",")
	}

	queryParams["show_alternative"] = fmt.Sprintf("%t", includeAlternativeRates)

	apiPath := timeSeriesAPIPath
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
	rsp := new(TimeSeriesResponse)

	// Because rsp is already a pointer
	if _, err = s.client.Do(req, rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}
