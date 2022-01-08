package responses

type Procurement struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Product                      string `json:"Product"`
			Plant                        string `json:"Plant"`
			IsAutoPurOrdCreationAllowed  bool   `json:"IsAutoPurOrdCreationAllowed"`
			IsSourceListRequired         bool   `json:"IsSourceListRequired"`
		} `json:"results"`
	} `json:"d"`
}
