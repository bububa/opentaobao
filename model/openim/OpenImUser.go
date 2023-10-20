package openim

import (
	"sync"
)

// OpenImUser 结构体
type OpenImUser struct {
	// 用户id
	Uid string `json:"uid,omitempty" xml:"uid,omitempty"`
	// 账户appkey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 是否为淘宝账号
	TaobaoAccount bool `json:"taobao_account,omitempty" xml:"taobao_account,omitempty"`
}

var poolOpenImUser = sync.Pool{
	New: func() any {
		return new(OpenImUser)
	},
}

// GetOpenImUser() 从对象池中获取OpenImUser
func GetOpenImUser() *OpenImUser {
	return poolOpenImUser.Get().(*OpenImUser)
}

// ReleaseOpenImUser 释放OpenImUser
func ReleaseOpenImUser(v *OpenImUser) {
	v.Uid = ""
	v.AppKey = ""
	v.TaobaoAccount = false
	poolOpenImUser.Put(v)
}
