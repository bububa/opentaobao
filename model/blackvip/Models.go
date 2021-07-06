package blackvip

// Models 结构体
type Models struct {
	// 88VIP到期时间
	VipExpireTime int64 `json:"vip_expire_time,omitempty" xml:"vip_expire_time,omitempty"`
	// 是否88VIP自动续费用户
	Is88VipAutoRenewUser bool `json:"is88_vip_auto_renew_user,omitempty" xml:"is88_vip_auto_renew_user,omitempty"`
	// 是否88VIP
	Is88Vip bool `json:"is88_vip,omitempty" xml:"is88_vip,omitempty"`
	// 是否88VIP潜在用户
	Is88VipTargetUser bool `json:"is88_vip_target_user,omitempty" xml:"is88_vip_target_user,omitempty"`
}
