package film

// LotteryPerformanceTopParam 结构体
type LotteryPerformanceTopParam struct {
	// 用户对外开放ID
	MixUserId string `json:"mix_user_id,omitempty" xml:"mix_user_id,omitempty"`
	// 活动ID
	SpreadId string `json:"spread_id,omitempty" xml:"spread_id,omitempty"`
	// 幂等ID
	OutBizId string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
	// 拓展参数
	BizExtInfo string `json:"biz_ext_info,omitempty" xml:"biz_ext_info,omitempty"`
	// 平台
	Platform int64 `json:"platform,omitempty" xml:"platform,omitempty"`
}
