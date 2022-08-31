package alihouse

// DynamicDataDto 结构体
type DynamicDataDto struct {
	// 小区outerid
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 业务时间
	BizTime string `json:"biz_time,omitempty" xml:"biz_time,omitempty"`
	// 政府限价
	GovernmentLimitPriceDesc string `json:"government_limit_price_desc,omitempty" xml:"government_limit_price_desc,omitempty"`
	// 月涨幅比（无百分号）
	MonthlyIncrease string `json:"monthly_increase,omitempty" xml:"monthly_increase,omitempty"`
	// 月均价
	MonthlyAvgPriceDesc string `json:"monthly_avg_price_desc,omitempty" xml:"monthly_avg_price_desc,omitempty"`
	// 月均价单位
	MonthlyAvgPriceUnit string `json:"monthly_avg_price_unit,omitempty" xml:"monthly_avg_price_unit,omitempty"`
}
