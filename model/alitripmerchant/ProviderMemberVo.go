package alitripmerchant

import (
	"sync"
)

// ProviderMemberVo 结构体
type ProviderMemberVo struct {
	// 卡类型
	CardType string `json:"card_type,omitempty" xml:"card_type,omitempty"`
	// 卡等级
	CardTier string `json:"card_tier,omitempty" xml:"card_tier,omitempty"`
	// 卡号
	CardNumber string `json:"card_number,omitempty" xml:"card_number,omitempty"`
	// 会员id
	MemberId string `json:"member_id,omitempty" xml:"member_id,omitempty"`
	// 注册类型
	RegisterType string `json:"register_type,omitempty" xml:"register_type,omitempty"`
}

var poolProviderMemberVo = sync.Pool{
	New: func() any {
		return new(ProviderMemberVo)
	},
}

// GetProviderMemberVo() 从对象池中获取ProviderMemberVo
func GetProviderMemberVo() *ProviderMemberVo {
	return poolProviderMemberVo.Get().(*ProviderMemberVo)
}

// ReleaseProviderMemberVo 释放ProviderMemberVo
func ReleaseProviderMemberVo(v *ProviderMemberVo) {
	v.CardType = ""
	v.CardTier = ""
	v.CardNumber = ""
	v.MemberId = ""
	v.RegisterType = ""
	poolProviderMemberVo.Put(v)
}
