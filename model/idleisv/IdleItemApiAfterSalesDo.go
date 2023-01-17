package idleisv

// IdleItemApiAfterSalesDo 结构体
type IdleItemApiAfterSalesDo struct {
	// 是否支持七天无理由
	SupportSdrPolicy bool `json:"support_sdr_policy,omitempty" xml:"support_sdr_policy,omitempty"`
	// 是否支持描述不符包邮退
	SupportNfrPolicy bool `json:"support_nfr_policy,omitempty" xml:"support_nfr_policy,omitempty"`
}
