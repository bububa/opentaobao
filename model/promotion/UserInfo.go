package promotion

import (
	"sync"
)

// UserInfo 结构体
type UserInfo struct {
	// 来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 用户名称
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 用户id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
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
	v.Source = ""
	v.UserName = ""
	v.UserId = 0
	poolUserInfo.Put(v)
}
