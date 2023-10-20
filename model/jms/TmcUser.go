package jms

import (
	"sync"
)

// TmcUser 结构体
type TmcUser struct {
	// 用户首次开通时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 用户最后开通时间
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 用户昵称
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 用户ID
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 用户授权是否有效，true表示授权有效，false表示授权过期
	IsValid bool `json:"is_valid,omitempty" xml:"is_valid,omitempty"`
}

var poolTmcUser = sync.Pool{
	New: func() any {
		return new(TmcUser)
	},
}

// GetTmcUser() 从对象池中获取TmcUser
func GetTmcUser() *TmcUser {
	return poolTmcUser.Get().(*TmcUser)
}

// ReleaseTmcUser 释放TmcUser
func ReleaseTmcUser(v *TmcUser) {
	v.Created = ""
	v.Modified = ""
	v.UserNick = ""
	v.UserId = 0
	v.IsValid = false
	poolTmcUser.Put(v)
}
