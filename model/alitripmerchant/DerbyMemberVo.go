package alitripmerchant

import (
	"sync"
)

// DerbyMemberVo 结构体
type DerbyMemberVo struct {
	// 用户臻享卡数量
	CardAmount string `json:"card_amount,omitempty" xml:"card_amount,omitempty"`
}

var poolDerbyMemberVo = sync.Pool{
	New: func() any {
		return new(DerbyMemberVo)
	},
}

// GetDerbyMemberVo() 从对象池中获取DerbyMemberVo
func GetDerbyMemberVo() *DerbyMemberVo {
	return poolDerbyMemberVo.Get().(*DerbyMemberVo)
}

// ReleaseDerbyMemberVo 释放DerbyMemberVo
func ReleaseDerbyMemberVo(v *DerbyMemberVo) {
	v.CardAmount = ""
	poolDerbyMemberVo.Put(v)
}
