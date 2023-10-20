package alilabs

import (
	"sync"
)

// BasicUserInfo 结构体
type BasicUserInfo struct {
	// 头像 url
	AvatarUrl string `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
	// 昵称
	NickName string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
}

var poolBasicUserInfo = sync.Pool{
	New: func() any {
		return new(BasicUserInfo)
	},
}

// GetBasicUserInfo() 从对象池中获取BasicUserInfo
func GetBasicUserInfo() *BasicUserInfo {
	return poolBasicUserInfo.Get().(*BasicUserInfo)
}

// ReleaseBasicUserInfo 释放BasicUserInfo
func ReleaseBasicUserInfo(v *BasicUserInfo) {
	v.AvatarUrl = ""
	v.NickName = ""
	poolBasicUserInfo.Put(v)
}
