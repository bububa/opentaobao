package alsc

import (
	"sync"
)

// UserInfo 结构体
type UserInfo struct {
	// 账号类型，0表示淘宝账号，25表示饿了么账号
	Site string `json:"site,omitempty" xml:"site,omitempty"`
	// 用户ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolUserInfo = sync.Pool{
	New: func() any {
		return new(UserInfo)
	},
}

// GetUserInfo() 从对象池中获取UserInfo
func GetUserInfo() *UserInfo {
	return poolUserInfo.Get().(*UserInfo)
}

// ReleaseUserInfo 释放UserInfo
func ReleaseUserInfo(v *UserInfo) {
	v.Site = ""
	v.UserId = ""
	poolUserInfo.Put(v)
}
