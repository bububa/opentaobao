package alimember

// SyncMemberDto 结构体
type SyncMemberDto struct {
	// 生日，格式yyyy-mm-dd
	Birthday string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 性别，F 女，M 男
	Sex string `json:"sex,omitempty" xml:"sex,omitempty"`
	// 手机号，海外手机格式 1-993333
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 外部会员ID，识别商家会员的唯一身份标识
	OuterMemberId string `json:"outer_member_id,omitempty" xml:"outer_member_id,omitempty"`
	// 用户ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 对外开放的merchantId
	OpenMerchantId string `json:"open_merchant_id,omitempty" xml:"open_merchant_id,omitempty"`
	// 用户站点，淘宝 taobao，饿了么 eleme，支付宝 alipay
	UserSite string `json:"user_site,omitempty" xml:"user_site,omitempty"`
	// 省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 外部会员卡号
	OuterCardNo string `json:"outer_card_no,omitempty" xml:"outer_card_no,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 授权账号类型，主账号/子账号
	UidType string `json:"uid_type,omitempty" xml:"uid_type,omitempty"`
	// 版本号
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 等级积分
	LevelPoint int64 `json:"level_point,omitempty" xml:"level_point,omitempty"`
	// 等级
	LevelNum int64 `json:"level_num,omitempty" xml:"level_num,omitempty"`
	// 消费积分
	ConsumePoint int64 `json:"consume_point,omitempty" xml:"consume_point,omitempty"`
}
