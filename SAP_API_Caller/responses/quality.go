package responses

type Quality struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Product                        string `json:"Product"`
			Plant                          string `json:"Plant"`
			MaximumStoragePeriod           string `json:"MaximumStoragePeriod"`
			QualityMgmtCtrlKey             string `json:"QualityMgmtCtrlKey"`
			MatlQualityAuthorizationGroup  string `json:"MatlQualityAuthorizationGroup"`
			HasPostToInspectionStock       bool   `json:"HasPostToInspectionStock"`
			InspLotDocumentationIsRequired bool   `json:"InspLotDocumentationIsRequired"`
			SuplrQualityManagementSystem   string `json:"SuplrQualityManagementSystem"`
			RecrrgInspIntervalTimeInDays   string `json:"RecrrgInspIntervalTimeInDays"`
			ProductQualityCertificateType  string `json:"ProductQualityCertificateType"`
		} `json:"results"`
	} `json:"d"`
}
