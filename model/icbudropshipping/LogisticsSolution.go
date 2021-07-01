package icbudropshipping

// LogisticsSolution 结构体
type LogisticsSolution struct {
	// delivery time (days)
	DeliveryTime string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
	// destination country
	DestinationCountry string `json:"destination_country,omitempty" xml:"destination_country,omitempty"`
	// dispatch country
	DispatchCountry string `json:"dispatch_country,omitempty" xml:"dispatch_country,omitempty"`
	// shipping fee
	Fee float64 `json:"fee,omitempty" xml:"fee,omitempty"`
	// shipping type
	ShippingType string `json:"shipping_type,omitempty" xml:"shipping_type,omitempty"`
	// trade term
	TradeTerm string `json:"trade_term,omitempty" xml:"trade_term,omitempty"`
	// vendor code
	VendorCode string `json:"vendor_code,omitempty" xml:"vendor_code,omitempty"`
	// vendor name
	VendorName string `json:"vendor_name,omitempty" xml:"vendor_name,omitempty"`
}
