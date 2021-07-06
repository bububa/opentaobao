package alsc

// JoinMemRuleOpenInfo 结构体
type JoinMemRuleOpenInfo struct {
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 是否删除
	Deleted bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// 关注公众号成为会员
	FollowWechatSupport bool `json:"follow_wechat_support,omitempty" xml:"follow_wechat_support,omitempty"`
	// 购买指定菜品成为会员
	JoinMemberDishSupport bool `json:"join_member_dish_support,omitempty" xml:"join_member_dish_support,omitempty"`
	// 手机号注册成为会员
	MobileRegisterSupport bool `json:"mobile_register_support,omitempty" xml:"mobile_register_support,omitempty"`
	// 支付成为会员
	PayJoinSupport bool `json:"pay_join_support,omitempty" xml:"pay_join_support,omitempty"`
	// 扫码点餐成为会员
	ScanOrderSupport bool `json:"scan_order_support,omitempty" xml:"scan_order_support,omitempty"`
}
