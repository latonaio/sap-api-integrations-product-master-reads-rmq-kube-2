package responses

type Plant struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Product                        string      `json:"Product"`
			Plant                          string      `json:"Plant"`
			PurchasingGroup                string      `json:"PurchasingGroup"`
			ProductionInvtryManagedLoc     string      `json:"ProductionInvtryManagedLoc"`
			AvailabilityCheckType          string      `json:"AvailabilityCheckType"`
			ProfitCenter                   string      `json:"ProfitCenter"`			
			GoodsReceiptDuration           string      `json:"GoodsReceiptDuration"`			
			MRPType                        string      `json:"MRPType"`
			MRPResponsible                 string      `json:"MRPResponsible"`			
			MinimumLotSizeQuantity         string      `json:"MinimumLotSizeQuantity"`
			MaximumLotSizeQuantity         string      `json:"MaximumLotSizeQuantity"`			
			FixedLotSizeQuantity           string      `json:"FixedLotSizeQuantity"`			
			IsBatchManagementRequired      bool        `json:"IsBatchManagementRequired"`
			ProcurementType                string      `json:"ProcurementType"`
			IsInternalBatchManaged         bool        `json:"IsInternalBatchManaged"`			
			GoodsIssueUnit                 string      `json:"GoodsIssueUnit"`	
			MaterialFreightGroup           string      `json:"MaterialFreightGroup"`
			ProductLogisticsHandlingGroup  string      `json:"ProductLogisticsHandlingGroup"`
			IsMarkedForDeletion            bool        `json:"IsMarkedForDeletion"`
		} `json:"results"`
	} `json:"d"`
}
