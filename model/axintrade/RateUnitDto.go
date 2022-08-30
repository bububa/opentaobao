package axintrade

// RateUnitDto 结构体
type RateUnitDto struct {
	// 房间价格
	DailyPriceInfoList []DailyPriceInfoDto `json:"daily_price_info_list,omitempty" xml:"daily_price_info_list>daily_price_info_dto,omitempty"`
}
