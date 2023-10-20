package alitripmerchant

import (
	"sync"
)

// LotteryDataVo 结构体
type LotteryDataVo struct {
	// 用户微信Open_ID
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 用户微信Union_ID
	UnionId string `json:"union_id,omitempty" xml:"union_id,omitempty"`
	// 雅高会员Pmid
	PmId string `json:"pm_id,omitempty" xml:"pm_id,omitempty"`
	// 会员卡号
	CardNumber string `json:"card_number,omitempty" xml:"card_number,omitempty"`
	// 注册时间
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 用户注册信息(注册渠道)
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
}

var poolLotteryDataVo = sync.Pool{
	New: func() any {
		return new(LotteryDataVo)
	},
}

// GetLotteryDataVo() 从对象池中获取LotteryDataVo
func GetLotteryDataVo() *LotteryDataVo {
	return poolLotteryDataVo.Get().(*LotteryDataVo)
}

// ReleaseLotteryDataVo 释放LotteryDataVo
func ReleaseLotteryDataVo(v *LotteryDataVo) {
	v.OpenId = ""
	v.UnionId = ""
	v.PmId = ""
	v.CardNumber = ""
	v.OperateTime = ""
	v.Remark = ""
	poolLotteryDataVo.Put(v)
}
