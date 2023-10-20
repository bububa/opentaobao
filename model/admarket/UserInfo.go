package admarket

import (
	"sync"
)

// UserInfo 结构体
type UserInfo struct {
	// 用户id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 补充信息
	Info string `json:"info,omitempty" xml:"info,omitempty"`
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
	v.Id = ""
	v.Info = ""
	poolUserInfo.Put(v)
}
