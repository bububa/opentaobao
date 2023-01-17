package caipiao

// LotteryWangcaiPresentStat 结构体
type LotteryWangcaiPresentStat struct {
	// 当日赠送彩票的注数
	PresentStake int64 `json:"present_stake,omitempty" xml:"present_stake,omitempty"`
	// 日期
	DateId int64 `json:"date_id,omitempty" xml:"date_id,omitempty"`
	// 当日赠送彩票的金额
	PresentFee int64 `json:"present_fee,omitempty" xml:"present_fee,omitempty"`
	// 送彩方的淘宝数字ID
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 当日赠送用户数
	PresentUser int64 `json:"present_user,omitempty" xml:"present_user,omitempty"`
}
