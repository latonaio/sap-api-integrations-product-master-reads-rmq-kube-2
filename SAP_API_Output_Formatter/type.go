package sap_api_output_formatter

type ProductMaster struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Deleted       string `json:"deleted"`
}

type General struct {
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
	ToProductDesc             string `json:"to_Description"`
}

type Plant struct {
	Product                       string `json:"Product"`
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
}

type MRPArea struct {
	Product                       string `json:"Product"`
	Plant                         string `json:"Plant"`
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
}

type WorkScheduling struct {
	Product                       string `json:"Product"`
	Plant                         string `json:"Plant"`
	ProductionInvtryManagedLoc    string `json:"ProductionInvtryManagedLoc"`
	ProductProcessingTime         string `json:"ProductProcessingTime"`
	ProductionSupervisor          string `json:"ProductionSupervisor"`
	ProductProductionQuantityUnit string `json:"ProductProductionQuantityUnit"`
	ProdnOrderIsBatchRequired     string `json:"ProdnOrderIsBatchRequired"`
	MatlCompIsMarkedForBackflush  string `json:"MatlCompIsMarkedForBackflush"`
	ProductionSchedulingProfile   string `json:"ProductionSchedulingProfile"`
}

type Procurement struct {
	Product                     string `json:"Product"`
	Plant                       string `json:"Plant"`
	IsAutoPurOrdCreationAllowed bool   `json:"IsAutoPurOrdCreationAllowed"`
	IsSourceListRequired        bool   `json:"IsSourceListRequired"`
}

type SalesPlant struct {
	Product               string `json:"Product"`
	Plant                 string `json:"Plant"`
	LoadingGroup          string `json:"LoadingGroup"`
	AvailabilityCheckType string `json:"AvailabilityCheckType"`
}

type Accounting struct {
	Product             string `json:"Product"`
	ValuationArea       string `json:"ValuationArea"`
	ValuationClass      string `json:"ValuationClass"`
	StandardPrice       string `json:"StandardPrice"`
	PriceUnitQty        string `json:"PriceUnitQty"`
	MovingAveragePrice  string `json:"MovingAveragePrice"`
	PriceLastChangeDate string `json:"PriceLastChangeDate"`
	PlannedPrice        string `json:"PlannedPrice"`
	IsMarkedForDeletion bool   `json:"IsMarkedForDeletion"`
}

type SalesOrganization struct {
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
}

type ProductDesc struct {
	Product            string `json:"Product"`
	Language           string `json:"Language"`
	ProductDescription string `json:"ProductDescription"`
}

type Quality struct {
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
}

type SalesTax struct {
	Product           string `json:"Product"`
	Country           string `json:"Country"`
	TaxCategory       string `json:"TaxCategory"`
	TaxClassification string `json:"TaxClassification"`
}

type ToProductDesc struct {
	Product            string `json:"Product"`
	Language           string `json:"Language"`
	ProductDescription string `json:"ProductDescription"`
}
