package wdk

// LimitDto 结构体
type LimitDto struct {
	// 赠品总限量/店
	TotalLimitCnt int64 `json:"total_limit_cnt,omitempty" xml:"total_limit_cnt,omitempty"`
	// 赠品每日总限量/店
	DailyTotalLimitCnt int64 `json:"daily_total_limit_cnt,omitempty" xml:"daily_total_limit_cnt,omitempty"`
	// 赠品每人限量/店
	UserTotalLimitCnt int64 `json:"user_total_limit_cnt,omitempty" xml:"user_total_limit_cnt,omitempty"`
	// 赠品每人每日限量/店
	UserDailyLimitCnt int64 `json:"user_daily_limit_cnt,omitempty" xml:"user_daily_limit_cnt,omitempty"`
	// 用户每单限购
	OrderLimitCnt int64 `json:"order_limit_cnt,omitempty" xml:"order_limit_cnt,omitempty"`
}
