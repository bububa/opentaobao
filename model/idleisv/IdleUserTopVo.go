package idleisv

// IdleUserTopVo 结构体
type IdleUserTopVo struct {
	// 用户选项信息: IDLE_UP（闲鱼升级用户），SDR（已开通七天无理由服务），NFR（已开通描述不符包邮退服务）
	Options []string `json:"options,omitempty" xml:"options>string,omitempty"`
	// 加密用户id
	EncryptUserid string `json:"encrypt_userid,omitempty" xml:"encrypt_userid,omitempty"`
	// 服务保障金信息
	ServiceDeposit *IdleDepositTopSubVo `json:"service_deposit,omitempty" xml:"service_deposit,omitempty"`
}
