package caipiao

import (
	"sync"
)

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

var poolLotteryWangcaiPresentStat = sync.Pool{
	New: func() any {
		return new(LotteryWangcaiPresentStat)
	},
}

// GetLotteryWangcaiPresentStat() 从对象池中获取LotteryWangcaiPresentStat
func GetLotteryWangcaiPresentStat() *LotteryWangcaiPresentStat {
	return poolLotteryWangcaiPresentStat.Get().(*LotteryWangcaiPresentStat)
}

// ReleaseLotteryWangcaiPresentStat 释放LotteryWangcaiPresentStat
func ReleaseLotteryWangcaiPresentStat(v *LotteryWangcaiPresentStat) {
	v.PresentStake = 0
	v.DateId = 0
	v.PresentFee = 0
	v.SellerId = 0
	v.PresentUser = 0
	poolLotteryWangcaiPresentStat.Put(v)
}
