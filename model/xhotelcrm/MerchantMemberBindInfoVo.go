package xhotelcrm

// MerchantMemberBindInfoVo 结构体
type MerchantMemberBindInfoVo struct {
	// 渠道
	Fpt string `json:"fpt,omitempty" xml:"fpt,omitempty"`
	// 注册时间
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 1.未注册/非支付宝端注册的会员2.未注册成功3.非新会员
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 是否完成任务
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
}
