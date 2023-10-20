package util

import (
	"sync"
)

// SessionInfo 结构体
type SessionInfo struct {
	// skey
	Skey string `json:"skey,omitempty" xml:"skey,omitempty"`
	// openId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// accessToken
	AccessToken string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	// unionId
	UnionId string `json:"union_id,omitempty" xml:"union_id,omitempty"`
}

var poolSessionInfo = sync.Pool{
	New: func() any {
		return new(SessionInfo)
	},
}

// GetSessionInfo() 从对象池中获取SessionInfo
func GetSessionInfo() *SessionInfo {
	return poolSessionInfo.Get().(*SessionInfo)
}

// ReleaseSessionInfo 释放SessionInfo
func ReleaseSessionInfo(v *SessionInfo) {
	v.Skey = ""
	v.OpenId = ""
	v.AccessToken = ""
	v.UnionId = ""
	poolSessionInfo.Put(v)
}
