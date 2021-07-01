package baodian

// CoinUserDepositV2 结构体
type CoinUserDepositV2 struct {
	// 欠款
	Credit int64 `json:"credit,omitempty" xml:"credit,omitempty"`
	// 欠款额度
	CreditLimit int64 `json:"credit_limit,omitempty" xml:"credit_limit,omitempty"`
	// 账期
	CreditPeriod string `json:"credit_period,omitempty" xml:"credit_period,omitempty"`
	// 宝点余额
	Deposit int64 `json:"deposit,omitempty" xml:"deposit,omitempty"`
	// 设备授权
	DeviceAuth bool `json:"device_auth,omitempty" xml:"device_auth,omitempty"`
	// 可以付款
	EnablePay bool `json:"enable_pay,omitempty" xml:"enable_pay,omitempty"`
	// 积分
	GamePoint int64 `json:"game_point,omitempty" xml:"game_point,omitempty"`
	// 图片地址
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 用户是否有逾期宝点贷款
	IsExpired bool `json:"is_expired,omitempty" xml:"is_expired,omitempty"`
	// 是否开通先用后付
	IsPayAfterPlay bool `json:"is_pay_after_play,omitempty" xml:"is_pay_after_play,omitempty"`
	// 是否新用户
	NewUser bool `json:"new_user,omitempty" xml:"new_user,omitempty"`
	// 单位数量(10个宝点)宝点价格，即10个宝点的价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// auth code
	UserAuthCode string `json:"user_auth_code,omitempty" xml:"user_auth_code,omitempty"`
	// 昵称
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 用户 str id
	UserStrId string `json:"user_str_id,omitempty" xml:"user_str_id,omitempty"`
	// 头像地址
	AvatarUrl string `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
}
