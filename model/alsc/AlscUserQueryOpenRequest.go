package alsc

import (
	"sync"
)

// AlscUserQueryOpenRequest 结构体
type AlscUserQueryOpenRequest struct {
	// 用户的授权令牌，通过请求授权令牌的接口获取。一个令牌对应一个用户的信息
	AccessToken string `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

var poolAlscUserQueryOpenRequest = sync.Pool{
	New: func() any {
		return new(AlscUserQueryOpenRequest)
	},
}

// GetAlscUserQueryOpenRequest() 从对象池中获取AlscUserQueryOpenRequest
func GetAlscUserQueryOpenRequest() *AlscUserQueryOpenRequest {
	return poolAlscUserQueryOpenRequest.Get().(*AlscUserQueryOpenRequest)
}

// ReleaseAlscUserQueryOpenRequest 释放AlscUserQueryOpenRequest
func ReleaseAlscUserQueryOpenRequest(v *AlscUserQueryOpenRequest) {
	v.AccessToken = ""
	poolAlscUserQueryOpenRequest.Put(v)
}
