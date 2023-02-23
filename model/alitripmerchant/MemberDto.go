package alitripmerchant

// MemberDto 结构体
type MemberDto struct {
	// 会员权益走马灯
	MemberRightCarousels []MemberRightCarouselVo `json:"member_right_carousels,omitempty" xml:"member_right_carousels>member_right_carousel_vo,omitempty"`
	// 用户id
	TenantUserId string `json:"tenant_user_id,omitempty" xml:"tenant_user_id,omitempty"`
	// 微信头像
	WechatAvatarUrl string `json:"wechat_avatar_url,omitempty" xml:"wechat_avatar_url,omitempty"`
	// 微信昵称
	WechatNickName string `json:"wechat_nick_name,omitempty" xml:"wechat_nick_name,omitempty"`
	// 手机号
	PhoneNum string `json:"phone_num,omitempty" xml:"phone_num,omitempty"`
	// 手机前缀
	PhonePri string `json:"phone_pri,omitempty" xml:"phone_pri,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 返回日期毫秒值
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 会员认证状态
	AuthorizeStatus string `json:"authorize_status,omitempty" xml:"authorize_status,omitempty"`
	// 会员注册状态
	RegisterType string `json:"register_type,omitempty" xml:"register_type,omitempty"`
	// 弹窗页面
	PopUpPage string `json:"pop_up_page,omitempty" xml:"pop_up_page,omitempty"`
	// 下一等级
	NextLevelName string `json:"next_level_name,omitempty" xml:"next_level_name,omitempty"`
	// 德比返回的手机号
	PhoneByDerby string `json:"phone_by_derby,omitempty" xml:"phone_by_derby,omitempty"`
	// 德比[查询用户信息]接口-[contactMedium]节点中的[id]
	DerbyEmailId string `json:"derby_email_id,omitempty" xml:"derby_email_id,omitempty"`
	// 用户基本信息
	MemberBaseInfo *MemberBaseInfoDto `json:"member_base_info,omitempty" xml:"member_base_info,omitempty"`
	// 会员卡信息
	CardBaseInfo *CardBaseInfoDto `json:"card_base_info,omitempty" xml:"card_base_info,omitempty"`
	// 当前用户拥有优惠券数量
	CouponCount int64 `json:"coupon_count,omitempty" xml:"coupon_count,omitempty"`
	// 会员卡信息
	MemberCardDetail *MemberCardDetailVo `json:"member_card_detail,omitempty" xml:"member_card_detail,omitempty"`
	// 有效代金券数量
	VoucherCount int64 `json:"voucher_count,omitempty" xml:"voucher_count,omitempty"`
	// 会员积分详情
	PointsDetail *PointsDetailDto `json:"points_detail,omitempty" xml:"points_detail,omitempty"`
	// 会员间夜详情
	NightDetail *PointsDetailDto `json:"night_detail,omitempty" xml:"night_detail,omitempty"`
	// 错误次数
	PwdErrorCount int64 `json:"pwd_error_count,omitempty" xml:"pwd_error_count,omitempty"`
	// 德比返回体中提取的参数
	DerbySpecialField *DerbySpecialField `json:"derby_special_field,omitempty" xml:"derby_special_field,omitempty"`
	// 是否是会员
	IsMember bool `json:"is_member,omitempty" xml:"is_member,omitempty"`
	// 是否有微信公众号openId
	HasWechatPublicAccountOpenId bool `json:"has_wechat_public_account_open_id,omitempty" xml:"has_wechat_public_account_open_id,omitempty"`
	// 是否为假邮箱
	IsFakeEmail bool `json:"is_fake_email,omitempty" xml:"is_fake_email,omitempty"`
	// 快速升级资格
	FastTrackEligibility bool `json:"fast_track_eligibility,omitempty" xml:"fast_track_eligibility,omitempty"`
}
