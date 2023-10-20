package subuser

import (
	"sync"
)

// SubUserDo 结构体
type SubUserDo struct {
	// 子账号用户名
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 主账号昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 子账号Id
	SubId int64 `json:"sub_id,omitempty" xml:"sub_id,omitempty"`
	// 子账号当前状态 1正常 -1删除  2冻结
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 子账号所属的主账号的唯一标识
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}

var poolSubUserDo = sync.Pool{
	New: func() any {
		return new(SubUserDo)
	},
}

// GetSubUserDo() 从对象池中获取SubUserDo
func GetSubUserDo() *SubUserDo {
	return poolSubUserDo.Get().(*SubUserDo)
}

// ReleaseSubUserDo 释放SubUserDo
func ReleaseSubUserDo(v *SubUserDo) {
	v.Nick = ""
	v.SellerNick = ""
	v.SubId = 0
	v.Status = 0
	v.SellerId = 0
	poolSubUserDo.Put(v)
}
