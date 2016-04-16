package model

// AutosuggestionResponse represents response from HERE Autosuggestion API
type AutosuggestionResponse struct {
	Suggestions []struct {
		Label       string `json:"label"`
		Language    string `json:"language"`
		CountryCode string `json:"countryCode"`
		LocationID  string `json:"locationId"`
		Address     struct {
			Country    string `json:"country"`
			State      string `json:"state"`
			County     string `json:"county"`
			City       string `json:"city"`
			District   string `json:"district"`
			Street     string `json:"street"`
			PostalCode string `json:"postalCode"`
		} `json:"address"`
		MatchLevel string `json:"matchLevel"`
	} `json:"suggestions"`
}
