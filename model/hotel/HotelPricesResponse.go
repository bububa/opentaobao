package hotel

import (
	"sync"
)

// HotelPricesResponse 结构体
type HotelPricesResponse struct {
	// 酒店报价列表
	HotelPrices []HotelPrice `json:"hotel_prices,omitempty" xml:"hotel_prices>hotel_price,omitempty"`
	// 用于请求复现
	SearchId string `json:"search_id,omitempty" xml:"search_id,omitempty"`
}

var poolHotelPricesResponse = sync.Pool{
	New: func() any {
		return new(HotelPricesResponse)
	},
}

// GetHotelPricesResponse() 从对象池中获取HotelPricesResponse
func GetHotelPricesResponse() *HotelPricesResponse {
	return poolHotelPricesResponse.Get().(*HotelPricesResponse)
}

// ReleaseHotelPricesResponse 释放HotelPricesResponse
func ReleaseHotelPricesResponse(v *HotelPricesResponse) {
	v.HotelPrices = v.HotelPrices[:0]
	v.SearchId = ""
	poolHotelPricesResponse.Put(v)
}
