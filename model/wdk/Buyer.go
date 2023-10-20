package wdk

import (
	"sync"
)

// Buyer 结构体
type Buyer struct {
	// 买家标识
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 买家用户名
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 买家昵称
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 买家电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 买家备注
	UserMemo string `json:"user_memo,omitempty" xml:"user_memo,omitempty"`
}

var poolBuyer = sync.Pool{
	New: func() any {
		return new(Buyer)
	},
}

// GetBuyer() 从对象池中获取Buyer
func GetBuyer() *Buyer {
	return poolBuyer.Get().(*Buyer)
}

// ReleaseBuyer 释放Buyer
func ReleaseBuyer(v *Buyer) {
	v.UserId = ""
	v.UserName = ""
	v.UserNick = ""
	v.Phone = ""
	v.UserMemo = ""
	poolBuyer.Put(v)
}
