package openim

import (
	"sync"
)

// User 结构体
type User struct {
	// 用户所属appkey
	Appkey string `json:"appkey,omitempty" xml:"appkey,omitempty"`
	// 用户id
	Uid string `json:"uid,omitempty" xml:"uid,omitempty"`
	// 账户appkey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 是否是淘宝账号
	TaobaoAccount bool `json:"taobao_account,omitempty" xml:"taobao_account,omitempty"`
}

var poolUser = sync.Pool{
	New: func() any {
		return new(User)
	},
}

// GetUser() 从对象池中获取User
func GetUser() *User {
	return poolUser.Get().(*User)
}

// ReleaseUser 释放User
func ReleaseUser(v *User) {
	v.Appkey = ""
	v.Uid = ""
	v.AppKey = ""
	v.TaobaoAccount = false
	poolUser.Put(v)
}
