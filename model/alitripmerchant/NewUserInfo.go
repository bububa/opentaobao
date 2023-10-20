package alitripmerchant

import (
	"sync"
)

// NewUserInfo 结构体
type NewUserInfo struct {
	// 国家
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 头像
	AvatarUrl string `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
	// 昵称
	NickName string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// 语言
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// 昵称
	NkName string `json:"nk_name,omitempty" xml:"nk_name,omitempty"`
	// 性别
	Gender int64 `json:"gender,omitempty" xml:"gender,omitempty"`
}

var poolNewUserInfo = sync.Pool{
	New: func() any {
		return new(NewUserInfo)
	},
}

// GetNewUserInfo() 从对象池中获取NewUserInfo
func GetNewUserInfo() *NewUserInfo {
	return poolNewUserInfo.Get().(*NewUserInfo)
}

// ReleaseNewUserInfo 释放NewUserInfo
func ReleaseNewUserInfo(v *NewUserInfo) {
	v.Country = ""
	v.Province = ""
	v.City = ""
	v.AvatarUrl = ""
	v.NickName = ""
	v.Language = ""
	v.NkName = ""
	v.Gender = 0
	poolNewUserInfo.Put(v)
}
