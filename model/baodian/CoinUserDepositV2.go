package baodian

import (
	"sync"
)

// CoinUserDepositV2 结构体
type CoinUserDepositV2 struct {
	// 账期
	CreditPeriod string `json:"credit_period,omitempty" xml:"credit_period,omitempty"`
	// 图片地址
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// auth code
	UserAuthCode string `json:"user_auth_code,omitempty" xml:"user_auth_code,omitempty"`
	// 昵称
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 用户 str id
	UserStrId string `json:"user_str_id,omitempty" xml:"user_str_id,omitempty"`
	// 头像地址
	AvatarUrl string `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
	// 欠款
	Credit int64 `json:"credit,omitempty" xml:"credit,omitempty"`
	// 欠款额度
	CreditLimit int64 `json:"credit_limit,omitempty" xml:"credit_limit,omitempty"`
	// 宝点余额
	Deposit int64 `json:"deposit,omitempty" xml:"deposit,omitempty"`
	// 积分
	GamePoint int64 `json:"game_point,omitempty" xml:"game_point,omitempty"`
	// 单位数量(10个宝点)宝点价格，即10个宝点的价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 设备授权
	DeviceAuth bool `json:"device_auth,omitempty" xml:"device_auth,omitempty"`
	// 可以付款
	EnablePay bool `json:"enable_pay,omitempty" xml:"enable_pay,omitempty"`
	// 用户是否有逾期宝点贷款
	IsExpired bool `json:"is_expired,omitempty" xml:"is_expired,omitempty"`
	// 是否开通先用后付
	IsPayAfterPlay bool `json:"is_pay_after_play,omitempty" xml:"is_pay_after_play,omitempty"`
	// 是否新用户
	NewUser bool `json:"new_user,omitempty" xml:"new_user,omitempty"`
}

var poolCoinUserDepositV2 = sync.Pool{
	New: func() any {
		return new(CoinUserDepositV2)
	},
}

// GetCoinUserDepositV2() 从对象池中获取CoinUserDepositV2
func GetCoinUserDepositV2() *CoinUserDepositV2 {
	return poolCoinUserDepositV2.Get().(*CoinUserDepositV2)
}

// ReleaseCoinUserDepositV2 释放CoinUserDepositV2
func ReleaseCoinUserDepositV2(v *CoinUserDepositV2) {
	v.CreditPeriod = ""
	v.ImageUrl = ""
	v.UserAuthCode = ""
	v.UserNick = ""
	v.UserStrId = ""
	v.AvatarUrl = ""
	v.Credit = 0
	v.CreditLimit = 0
	v.Deposit = 0
	v.GamePoint = 0
	v.Price = 0
	v.DeviceAuth = false
	v.EnablePay = false
	v.IsExpired = false
	v.IsPayAfterPlay = false
	v.NewUser = false
	poolCoinUserDepositV2.Put(v)
}
