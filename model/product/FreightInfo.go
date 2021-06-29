package product

// FreightInfo 
type FreightInfo struct {
    // delivery_time
    DeliveryTime   string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
    // logistic_company
    LogisticCompany   string `json:"logistic_company,omitempty" xml:"logistic_company,omitempty"`
    // freight_cost_cent
    FreightCostCent   int64 `json:"freight_cost_cent,omitempty" xml:"freight_cost_cent,omitempty"`
    // freight_cost_currency
    FreightCostCurrency   string `json:"freight_cost_currency,omitempty" xml:"freight_cost_currency,omitempty"`
}
