package hotel

// HotelPricesResponse 结构体
type HotelPricesResponse struct {
	// 酒店报价列表
	HotelPrices []HotelPrice `json:"hotel_prices,omitempty" xml:"hotel_prices>hotel_price,omitempty"`
	// 用于请求复现
	SearchId string `json:"search_id,omitempty" xml:"search_id,omitempty"`
}
