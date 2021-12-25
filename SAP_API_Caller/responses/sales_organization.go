package responses

type SalesOrganization struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Product                        string `json:"Product"`
			ProductSalesOrg                string `json:"ProductSalesOrg"`
			ProductDistributionChnl        string `json:"ProductDistributionChnl"`
			SupplyingPlant                 string `json:"SupplyingPlant"`
			PriceSpecificationProductGroup string `json:"PriceSpecificationProductGroup"`
			AccountDetnProductGroup        string `json:"AccountDetnProductGroup"`
			ItemCategoryGroup              string `json:"ItemCategoryGroup"`
			SalesMeasureUnit               string `json:"SalesMeasureUnit"`
			ProductHierarchy               string `json:"ProductHierarchy"`
			IsMarkedForDeletion            bool   `json:"IsMarkedForDeletion"`
		} `json:"results"`
	} `json:"d"`
}
