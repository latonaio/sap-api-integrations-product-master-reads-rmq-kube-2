package responses

type SalesPlant struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Product               string `json:"Product"`
			Plant                 string `json:"Plant"`
			LoadingGroup          string `json:"LoadingGroup"`
			AvailabilityCheckType string `json:"AvailabilityCheckType"`
		} `json:"results"`
	} `json:"d"`
}
