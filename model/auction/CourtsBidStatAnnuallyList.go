package auction

// CourtsBidStatAnnuallyList 结构体
type CourtsBidStatAnnuallyList struct {
	// 时间区间(年)
	Period string `json:"period,omitempty" xml:"period,omitempty"`
	// 法院按年维度标的统计汇总
	CourtsBidStatSum *CourtsBidStatSum `json:"courts_bid_stat_sum,omitempty" xml:"courts_bid_stat_sum,omitempty"`
}
