package auction

// CourtsBidStatMonthlyList 结构体
type CourtsBidStatMonthlyList struct {
	// 法院按月维度标的统计汇总
	CourtsBidStatSum *CourtsBidStatSum `json:"courts_bid_stat_sum,omitempty" xml:"courts_bid_stat_sum,omitempty"`
	// 时间区间(月份)
	Period string `json:"period,omitempty" xml:"period,omitempty"`
}
