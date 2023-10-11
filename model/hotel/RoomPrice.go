package hotel

// RoomPrice 结构体
type RoomPrice struct {
	// 报价列表
	RatePrices []RatePrice `json:"rate_prices,omitempty" xml:"rate_prices>rate_price,omitempty"`
	// 房型名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 床型信息json
	BedJson string `json:"bed_json,omitempty" xml:"bed_json,omitempty"`
	// 窗型信息json 0无窗 1有窗 2部分有窗 3暗窗 4部分有窗 5落地窗
	WindowJson string `json:"window_json,omitempty" xml:"window_json,omitempty"`
	// 房型id
	Srid int64 `json:"srid,omitempty" xml:"srid,omitempty"`
}
