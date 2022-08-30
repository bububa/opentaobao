package idleisv

// IdleUserApiDo 结构体
type IdleUserApiDo struct {
	// 具备准入权限的业务类型列表
	SupportBizTypes []SupportBizType `json:"support_biz_types,omitempty" xml:"support_biz_types>support_biz_type,omitempty"`
	// 具备准入权限的行业品类列表
	SupportCatTypes []SupportCatType `json:"support_cat_types,omitempty" xml:"support_cat_types>support_cat_type,omitempty"`
	// 卖家的snsNick（已经废弃，不再返回数据）
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 用户身份（GENERAL: 普通用户，PRO_PLAYER: 个人经营者）
	Identity string `json:"identity,omitempty" xml:"identity,omitempty"`
	// 加密的卖家userId
	EncryptionSellerId string `json:"encryption_seller_id,omitempty" xml:"encryption_seller_id,omitempty"`
	// 是否是账号独立升级用户，是否有闲鱼独立账号升级标
	HasUptag bool `json:"has_uptag,omitempty" xml:"has_uptag,omitempty"`
}
