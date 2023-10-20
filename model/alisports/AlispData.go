package alisports

import (
	"sync"
)

// AlispData 结构体
type AlispData struct {
	// sso_token
	SsoToken string `json:"sso_token,omitempty" xml:"sso_token,omitempty"`
	// aliuid
	Aliuid string `json:"aliuid,omitempty" xml:"aliuid,omitempty"`
	// access_token
	AccessToken string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	// 头像
	AvatarUrl string `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
	// 昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 手机号，需要找阿里体育申请并且配置白名单返回
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// type
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolAlispData = sync.Pool{
	New: func() any {
		return new(AlispData)
	},
}

// GetAlispData() 从对象池中获取AlispData
func GetAlispData() *AlispData {
	return poolAlispData.Get().(*AlispData)
}

// ReleaseAlispData 释放AlispData
func ReleaseAlispData(v *AlispData) {
	v.SsoToken = ""
	v.Aliuid = ""
	v.AccessToken = ""
	v.AvatarUrl = ""
	v.Nick = ""
	v.Mobile = ""
	v.Type = 0
	poolAlispData.Put(v)
}
