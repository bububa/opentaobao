package tbk

// ServiceFeeDto 结构体
type ServiceFeeDto struct {
	// 专项服务费率
	ShareRelativeRate string `json:"share_relative_rate,omitempty" xml:"share_relative_rate,omitempty"`
	// 结算专项服务费
	ShareFee string `json:"share_fee,omitempty" xml:"share_fee,omitempty"`
	// 预估专项服务费
	SharePreFee string `json:"share_pre_fee,omitempty" xml:"share_pre_fee,omitempty"`
	// 专项服务费来源，122-渠道
	TkShareRoleType int64 `json:"tk_share_role_type,omitempty" xml:"tk_share_role_type,omitempty"`
}
