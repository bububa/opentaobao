package idleisv

// IdleItemApiAfterSalesDo 结构体
type IdleItemApiAfterSalesDo struct {
	// 是否支持七天无理由
	SupportSdrPolicy bool `json:"support_sdr_policy,omitempty" xml:"support_sdr_policy,omitempty"`
	// 是否支持描述不符包邮退
	SupportNfrPolicy bool `json:"support_nfr_policy,omitempty" xml:"support_nfr_policy,omitempty"`
	// 是否支持 虚拟-描述不符包退
	SupportVnrPolicy bool `json:"support_vnr_policy,omitempty" xml:"support_vnr_policy,omitempty"`
	// 是否支持 极速发货-24小时
	SupportFd24hsPolicy bool `json:"support_fd24hs_policy,omitempty" xml:"support_fd24hs_policy,omitempty"`
	// 是否支持 极速发货-10分钟
	SupportFd10msPolicy bool `json:"support_fd10ms_policy,omitempty" xml:"support_fd10ms_policy,omitempty"`
}
