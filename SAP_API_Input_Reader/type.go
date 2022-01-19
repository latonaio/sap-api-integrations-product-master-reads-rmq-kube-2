package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product      struct {
		Product             string `json:"Material"`
		IndustrySector      string `json:"IndustrySector"`
		ProductType         string `json:"ProductType"`
		BaseUnit            string `json:"BaseUnit"`
		ValidityStartDate   string `json:"ValidityStartDate"`
		ProductGroup        string `json:"ProductGroup"`
		Division            string `json:"Division"`
		GrossWeight         string `json:"GrossWeight"`
		WeightUnit          string `json:"WeightUnit"`
		SizeOrDimensionText string `json:"SizeOrDimensionText"`
		ProductStandardID   string `json:"ProductStandardID"`
		CreationDate        string `json:"CreationDate"`
		LastChangeDate      string `json:"LastChangeDate"`
		IsMarkedForDeletion string `json:"IsMarkedForDeletion"`
		NetWeight           string `json:"NetWeight"`
		ChangeNumber        string `json:"ChangeNumber"`
		Plant               struct {
			Plant                         string `json:"Plant"`
			PurchasingGroup               string `json:"PurchasingGroup"`
			ProductionInvtryManagedLoc    string `json:"ProductionInvtryManagedLoc"`
			AvailabilityCheckType         string `json:"AvailabilityCheckType"`
			ProfitCenter                  string `json:"ProfitCenter"`
			GoodsReceiptDuration          string `json:"GoodsReceiptDuration"`
			MRPType                       string `json:"MRPType"`
			MRPResponsible                string `json:"MRPResponsible"`
			MinimumLotSizeQuantity        string `json:"MinimumLotSizeQuantity"`
			MaximumLotSizeQuantity        string `json:"MaximumLotSizeQuantity"`
			FixedLotSizeQuantity          string `json:"FixedLotSizeQuantity"`
			IsBatchManagementRequired     bool   `json:"IsBatchManagementRequired"`
			ProcurementType               string `json:"ProcurementType"`
			IsInternalBatchManaged        bool   `json:"IsInternalBatchManaged"`
			GoodsIssueUnit                string `json:"GoodsIssueUnit"`
			MaterialFreightGroup          string `json:"MaterialFreightGroup"`
			ProductLogisticsHandlingGroup string `json:"ProductLogisticsHandlingGroup"`
			IsMarkedForDeletion           bool   `json:"IsMarkedForDeletion"`
			MRPArea                       struct {
				MRPArea                       string `json:"MRPArea"`
				MRPType                       string `json:"MRPType"`
				MRPResponsible                string `json:"MRPResponsible"`
				MRPGroup                      string `json:"MRPGroup"`
				ReorderThresholdQuantity      string `json:"ReorderThresholdQuantity"`
				PlanningTimeFence             string `json:"PlanningTimeFence"`
				LotSizingProcedure            string `json:"LotSizingProcedure"`
				LotSizeRoundingQuantity       string `json:"LotSizeRoundingQuantity"`
				MinimumLotSizeQuantity        string `json:"MinimumLotSizeQuantity"`
				MaximumLotSizeQuantity        string `json:"MaximumLotSizeQuantity"`
				MaximumStockQuantity          string `json:"MaximumStockQuantity"`
				ProcurementSubType            string `json:"ProcurementSubType"`
				DfltStorageLocationExtProcmt  string `json:"DfltStorageLocationExtProcmt"`
				MRPPlanningCalendar           string `json:"MRPPlanningCalendar"`
				SafetyStockQuantity           string `json:"SafetyStockQuantity"`
				SafetyDuration                string `json:"SafetyDuration"`
				FixedLotSizeQuantity          string `json:"FixedLotSizeQuantity"`
				PlannedDeliveryDurationInDays string `json:"PlannedDeliveryDurationInDays"`
				StorageLocation               string `json:"StorageLocation"`
				IsMarkedForDeletion           bool   `json:"IsMarkedForDeletion"`
			} `json:"MRPArea"`
			Procurement struct {
				IsAutoPurOrdCreationAllowed bool `json:"IsAutoPurOrdCreationAllowed"`
				IsSourceListRequired        bool `json:"IsSourceListRequired"`
			} `json:"Procurement"`
			SalesPlant struct {
				LoadingGroup          string `json:"LoadingGroup"`
				AvailabilityCheckType string `json:"AvailabilityCheckType"`
			} `json:"Sales_Plant"`
			WorkScheduling struct {
				ProductionInvtryManagedLoc    string `json:"ProductionInvtryManagedLoc"`
				ProductProcessingTime         string `json:"ProductProcessingTime"`
				ProductionSupervisor          string `json:"ProductionSupervisor"`
				ProductProductionQuantityUnit string `json:"ProductProductionQuantityUnit"`
				ProdnOrderIsBatchRequired     string `json:"ProdnOrderIsBatchRequired"`
				MatlCompIsMarkedForBackflush  string `json:"MatlCompIsMarkedForBackflush"`
				ProductionSchedulingProfile   string `json:"ProductionSchedulingProfile"`
			} `json:"Work_Scheduling"`
			Quality struct {
				MaximumStoragePeriod           string      `json:"MaximumStoragePeriod"`
				QualityMgmtCtrlKey             string      `json:"QualityMgmtCtrlKey"`
				MatlQualityAuthorizationGroup  string      `json:"MatlQualityAuthorizationGroup"`
				HasPostToInspectionStock       bool        `json:"HasPostToInspectionStock"`
				InspLotDocumentationIsRequired bool        `json:"InspLotDocumentationIsRequired"`
				SuplrQualityManagementSystem   string      `json:"SuplrQualityManagementSystem"`
				RecrrgInspIntervalTimeInDays   string      `json:"RecrrgInspIntervalTimeInDays"`
				ProductQualityCertificateType  string      `json:"ProductQualityCertificateType"`
			} `json:"Quality"`
		} `json:"Plant"`
		SalesOrganization struct {
			ProductSalesOrg                string `json:"ProductSalesOrg"`
			ProductDistributionChnl        string `json:"ProductDistributionChnl"`
			SupplyingPlant                 string `json:"SupplyingPlant"`
			PriceSpecificationProductGroup string `json:"PriceSpecificationProductGroup"`
			AccountDetnProductGroup        string `json:"AccountDetnProductGroup"`
			ItemCategoryGroup              string `json:"ItemCategoryGroup"`
			SalesMeasureUnit               string `json:"SalesMeasureUnit"`
			ProductHierarchy               string `json:"ProductHierarchy"`
			IsMarkedForDeletion            bool   `json:"IsMarkedForDeletion"`
		} `json:"Sales_Organization"`
		Accounting struct {
			ValuationArea       string `json:"ValuationArea"`
			ValuationClass      string `json:"ValuationClass"`
			StandardPrice       string `json:"StandardPrice"`
			PriceUnitQty        string `json:"PriceUnitQty"`
			MovingAveragePrice  string `json:"MovingAveragePrice"`
			PriceLastChangeDate string `json:"PriceLastChangeDate"`
			PlannedPrice        string `json:"PlannedPrice"`
			IsMarkedForDeletion bool   `json:"IsMarkedForDeletion"`
		} `json:"ValuationArea"`
		ProductDescription struct {
			Language           string `json:"Language"`
			ProductDescription string `json:"ProductDescription"`
		} `json:"ProductDescription"`
		SalesTax struct {
			Country           string `json:"Country"`
			TaxCategory       string `json:"TaxCategory"`
			TaxClassification string `json:"TaxClassification"`
		} `json:"SalesTax"`
	} `json:"material"`
	APISchema    string   `json:"api_schema"`
	Accepter     []string `json:"accepter"`
	MaterialCode string   `json:"material_code"`
	Deleted      bool     `json:"deleted"`
}
