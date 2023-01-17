package flight

// SaleDto 结构体
type SaleDto struct {
	// 销售日期
	SaleDate []DatePairDto `json:"sale_date,omitempty" xml:"sale_date>date_pair_dto,omitempty"`
	// 0-B2C, 1-B2B, 2-B2G, 3-宝贝。可多选，多个请用英文”,”隔开
	GoodsMarket []string `json:"goods_market,omitempty" xml:"goods_market>string,omitempty"`
	// 销售日期
	SaleDates []DatePairDto `json:"sale_dates,omitempty" xml:"sale_dates>date_pair_dto,omitempty"`
	// 提前预定天数
	AdvanceDay string `json:"advance_day,omitempty" xml:"advance_day,omitempty"`
	// 销售时间
	SaleTime string `json:"sale_time,omitempty" xml:"sale_time,omitempty"`
	// 只能输入A-B的格式，AB都是0-365之间的整数。往返航班之间的停留天数，以天数为单位，默认不限制。空表示默认
	StayDay string `json:"stay_day,omitempty" xml:"stay_day,omitempty"`
	// 销售方式：0，无；1，打包销售套餐1；2，打包销售套餐2；3，打包销售套餐3；4，返现-航司运价；5，返现-销售方包装；6，花呗卖家版；
	SaleModeCode int64 `json:"sale_mode_code,omitempty" xml:"sale_mode_code,omitempty"`
}
