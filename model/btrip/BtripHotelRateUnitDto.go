package btrip

// BtripHotelRateUnitDto 结构体
type BtripHotelRateUnitDto struct {
	// 日历信息
	DailyPriceInfoList []BtripHotelDailyPriceInfoDto `json:"daily_price_info_list,omitempty" xml:"daily_price_info_list>btrip_hotel_daily_price_info_dto,omitempty"`
	// 最小售卖单元唯一key
	RateKey string `json:"rate_key,omitempty" xml:"rate_key,omitempty"`
}
