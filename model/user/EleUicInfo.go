package user

import (
	"sync"
)

// EleUicInfo 结构体
type EleUicInfo struct {
	// 用户头像
	Avatar string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 用户昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 手机号
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// openId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
}

var poolEleUicInfo = sync.Pool{
	New: func() any {
		return new(EleUicInfo)
	},
}

// GetEleUicInfo() 从对象池中获取EleUicInfo
func GetEleUicInfo() *EleUicInfo {
	return poolEleUicInfo.Get().(*EleUicInfo)
}

// ReleaseEleUicInfo 释放EleUicInfo
func ReleaseEleUicInfo(v *EleUicInfo) {
	v.Avatar = ""
	v.Nick = ""
	v.Phone = ""
	v.OpenId = ""
	poolEleUicInfo.Put(v)
}
