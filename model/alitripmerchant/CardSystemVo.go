package alitripmerchant

import (
	"sync"
)

// CardSystemVo 结构体
type CardSystemVo struct {
	// 会员卡详细信息
	MemberCardDetailList []MemberCardDetailVo `json:"member_card_detail_list,omitempty" xml:"member_card_detail_list>member_card_detail_vo,omitempty"`
}

var poolCardSystemVo = sync.Pool{
	New: func() any {
		return new(CardSystemVo)
	},
}

// GetCardSystemVo() 从对象池中获取CardSystemVo
func GetCardSystemVo() *CardSystemVo {
	return poolCardSystemVo.Get().(*CardSystemVo)
}

// ReleaseCardSystemVo 释放CardSystemVo
func ReleaseCardSystemVo(v *CardSystemVo) {
	v.MemberCardDetailList = v.MemberCardDetailList[:0]
	poolCardSystemVo.Put(v)
}
