package aedropshiper

// LogisticsInfoDto 结构体
type LogisticsInfoDto struct {
	// Country
	ShipToCountry string `json:"ship_to_country,omitempty" xml:"ship_to_country,omitempty"`
	// Goods lead time
	DeliveryTime int64 `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
}
