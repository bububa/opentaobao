package tmallgenie

// VipCycleThirdVipStatusRequest 结构体
type VipCycleThirdVipStatusRequest struct {
	// 三方用户id
	ThirdUserId string `json:"third_user_id,omitempty" xml:"third_user_id,omitempty"`
	// 三方源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 合约类型：1-包月 2-包季 3-包年
	PerformType int64 `json:"perform_type,omitempty" xml:"perform_type,omitempty"`
}
