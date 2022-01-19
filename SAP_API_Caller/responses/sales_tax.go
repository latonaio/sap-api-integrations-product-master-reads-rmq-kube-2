package responses

type SalesTax struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Product           string `json:"Product"`
			Country           string `json:"Country"`
			TaxCategory       string `json:"TaxCategory"`
			TaxClassification string `json:"TaxClassification"`
		} `json:"results"`
	} `json:"d"`
}
