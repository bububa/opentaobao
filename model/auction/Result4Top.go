package auction

// Result4Top 结构体
type Result4Top struct {
	// 最新出价列表
	Results []LatestBids `json:"results,omitempty" xml:"results>latest_bids,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误说明
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回的最新出价条数
	TotalItem int64 `json:"total_item,omitempty" xml:"total_item,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
