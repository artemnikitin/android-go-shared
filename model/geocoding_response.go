package model

// GeocodingResponse contains response from HERE Geocoding API
type GeocodingResponse struct {
	Response struct {
		MetaInfo struct {
			Timestamp string `json:"Timestamp"`
		} `json:"MetaInfo"`
		View []struct {
			Result []struct {
				Location struct {
					Address struct {
						AdditionalData []struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"AdditionalData"`
						City        string `json:"City"`
						Country     string `json:"Country"`
						County      string `json:"County"`
						District    string `json:"District"`
						HouseNumber string `json:"HouseNumber"`
						Label       string `json:"Label"`
						PostalCode  string `json:"PostalCode"`
						State       string `json:"State"`
						Street      string `json:"Street"`
					} `json:"Address"`
					DisplayPosition struct {
						Latitude  float64 `json:"Latitude"`
						Longitude float64 `json:"Longitude"`
					} `json:"DisplayPosition"`
					LocationID   string `json:"LocationId"`
					LocationType string `json:"LocationType"`
					MapView      struct {
						BottomRight struct {
							Latitude  float64 `json:"Latitude"`
							Longitude float64 `json:"Longitude"`
						} `json:"BottomRight"`
						TopLeft struct {
							Latitude  float64 `json:"Latitude"`
							Longitude float64 `json:"Longitude"`
						} `json:"TopLeft"`
					} `json:"MapView"`
					NavigationPosition []struct {
						Latitude  float64 `json:"Latitude"`
						Longitude float64 `json:"Longitude"`
					} `json:"NavigationPosition"`
				} `json:"Location"`
				MatchLevel   string `json:"MatchLevel"`
				MatchQuality struct {
					City        float64   `json:"City"`
					HouseNumber float64   `json:"HouseNumber"`
					Street      []float64 `json:"Street"`
				} `json:"MatchQuality"`
				MatchType string  `json:"MatchType"`
				Relevance float64 `json:"Relevance"`
			} `json:"Result"`
			ViewID int    `json:"ViewId"`
			Type   string `json:"_type"`
		} `json:"View"`
	} `json:"Response"`
}
