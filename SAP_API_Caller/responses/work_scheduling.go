package responses

type WorkScheduling struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Product                       string `json:"Product"`
			Plant                         string `json:"Plant"`
			ProductionInvtryManagedLoc    string `json:"ProductionInvtryManagedLoc"`
			ProductProcessingTime         string `json:"ProductProcessingTime"`
			ProductionSupervisor          string `json:"ProductionSupervisor"`
			ProductProductionQuantityUnit string `json:"ProductProductionQuantityUnit"`
			ProdnOrderIsBatchRequired     string `json:"ProdnOrderIsBatchRequired"`
			MatlCompIsMarkedForBackflush  string `json:"MatlCompIsMarkedForBackflush"`
			ProductionSchedulingProfile   string `json:"ProductionSchedulingProfile"`
		} `json:"results"`
	} `json:"d"`
}
