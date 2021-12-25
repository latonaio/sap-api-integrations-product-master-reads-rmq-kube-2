package responses

type ProductDesc struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Product            string `json:"Product"`
			Language           string `json:"Language"`
			ProductDescription string `json:"ProductDescription"`
		} `json:"results"`
	} `json:"d"`
}
