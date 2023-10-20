package game

import (
	"sync"
)

// AppleTopNewSignNotifyDo 结构体
type AppleTopNewSignNotifyDo struct {
	// 电子卡卡号
	CardNo string `json:"card_no,omitempty" xml:"card_no,omitempty"`
	// TopUpTocken
	Tut string `json:"tut,omitempty" xml:"tut,omitempty"`
	// 商户订单号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 附加信息，后续可以扩展
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 商户上送的用户Id
	UserNo string `json:"user_no,omitempty" xml:"user_no,omitempty"`
	// 商户上送UCI
	Uci string `json:"uci,omitempty" xml:"uci,omitempty"`
	// Mask id
	Mai string `json:"mai,omitempty" xml:"mai,omitempty"`
	// 网关请求流水号,每次请求唯一
	QueryId string `json:"query_id,omitempty" xml:"query_id,omitempty"`
	// 商户上送tUCI
	Tuci string `json:"tuci,omitempty" xml:"tuci,omitempty"`
}

var poolAppleTopNewSignNotifyDo = sync.Pool{
	New: func() any {
		return new(AppleTopNewSignNotifyDo)
	},
}

// GetAppleTopNewSignNotifyDo() 从对象池中获取AppleTopNewSignNotifyDo
func GetAppleTopNewSignNotifyDo() *AppleTopNewSignNotifyDo {
	return poolAppleTopNewSignNotifyDo.Get().(*AppleTopNewSignNotifyDo)
}

// ReleaseAppleTopNewSignNotifyDo 释放AppleTopNewSignNotifyDo
func ReleaseAppleTopNewSignNotifyDo(v *AppleTopNewSignNotifyDo) {
	v.CardNo = ""
	v.Tut = ""
	v.OrderNo = ""
	v.Memo = ""
	v.UserNo = ""
	v.Uci = ""
	v.Mai = ""
	v.QueryId = ""
	v.Tuci = ""
	poolAppleTopNewSignNotifyDo.Put(v)
}
