package flight

// FlightPriceDto 结构体
type FlightPriceDto struct {
	// 指定舱位
	FareCabin string `json:"fare_cabin,omitempty" xml:"fare_cabin,omitempty"`
	// 行程索引
	FlightIndex string `json:"flight_index,omitempty" xml:"flight_index,omitempty"`
	// 指定的fareBasis
	MatchFareBasis string `json:"match_fare_basis,omitempty" xml:"match_fare_basis,omitempty"`
	// 直减金额
	DownFare int64 `json:"down_fare,omitempty" xml:"down_fare,omitempty"`
	// 直减比例
	DownPercent int64 `json:"down_percent,omitempty" xml:"down_percent,omitempty"`
	// 验舱
	ValidateCabin int64 `json:"validate_cabin,omitempty" xml:"validate_cabin,omitempty"`
}
