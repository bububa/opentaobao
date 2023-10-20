package tbk

// ServiceFeeDto 结构体
type ServiceFeeDto struct {
	// 专项服务费率（字段已废弃）
	Sharerelativerate string `json:"share_relative_rate,omitempty" xml:"share_relative_rate,omitempty"`
	// 结算专项服务费（字段已废弃）
	Sharefee string `json:"share_fee,omitempty" xml:"share_fee,omitempty"`
	// 预估专项服务费（字段已废弃）
	Shareprefee string `json:"share_pre_fee,omitempty" xml:"share_pre_fee,omitempty"`
	// 专项服务费来源，122-渠道（字段已废弃）
	Tkshareroletype int64 `json:"tk_share_role_type,omitempty" xml:"tk_share_role_type,omitempty"`
}
