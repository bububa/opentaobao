package openim

import (
	"sync"
)

// TribeUser 结构体
type TribeUser struct {
	// 用户id
	Uid string `json:"uid,omitempty" xml:"uid,omitempty"`
	// 群成员角色
	Role string `json:"role,omitempty" xml:"role,omitempty"`
	// 账户appkey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 是否为淘宝账号
	TaobaoAccount bool `json:"taobao_account,omitempty" xml:"taobao_account,omitempty"`
}

var poolTribeUser = sync.Pool{
	New: func() any {
		return new(TribeUser)
	},
}

// GetTribeUser() 从对象池中获取TribeUser
func GetTribeUser() *TribeUser {
	return poolTribeUser.Get().(*TribeUser)
}

// ReleaseTribeUser 释放TribeUser
func ReleaseTribeUser(v *TribeUser) {
	v.Uid = ""
	v.Role = ""
	v.AppKey = ""
	v.TaobaoAccount = false
	poolTribeUser.Put(v)
}
