package responses

type General struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Product                   string `json:"Product"`
			IndustrySector            string `json:"IndustrySector"`
			ProductType               string `json:"ProductType"`
			BaseUnit                  string `json:"BaseUnit"`
			ValidityStartDate         string `json:"ValidityStartDate"`
			ProductGroup              string `json:"ProductGroup"`
			Division                  string `json:"Division"`
			GrossWeight               string `json:"GrossWeight"`
			WeightUnit                string `json:"WeightUnit"`
			SizeOrDimensionText       string `json:"SizeOrDimensionText"`
			ProductStandardID         string `json:"ProductStandardID"`
			CreationDate              string `json:"CreationDate"`
			LastChangeDate            string `json:"LastChangeDate"`
			IsMarkedForDeletion       bool   `json:"IsMarkedForDeletion"`
			NetWeight                 string `json:"NetWeight"`
			ChangeNumber              string `json:"ChangeNumber"`
			ToProductDesc             struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Description"`
		} `json:"results"`
	} `json:"d"`
}
