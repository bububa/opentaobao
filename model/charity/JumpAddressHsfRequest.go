package charity

import (
	"sync"
)

// JumpAddressHsfRequest 结构体
type JumpAddressHsfRequest struct {
	// 三方用户昵称，建议脱敏
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// appKey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 三方用户id
	UserKey string `json:"user_key,omitempty" xml:"user_key,omitempty"`
	// 跳转平台的类型，会生成不同平台的uri ALIPAY：支付宝 OTHER:其他
	Platform string `json:"platform,omitempty" xml:"platform,omitempty"`
}

var poolJumpAddressHsfRequest = sync.Pool{
	New: func() any {
		return new(JumpAddressHsfRequest)
	},
}

// GetJumpAddressHsfRequest() 从对象池中获取JumpAddressHsfRequest
func GetJumpAddressHsfRequest() *JumpAddressHsfRequest {
	return poolJumpAddressHsfRequest.Get().(*JumpAddressHsfRequest)
}

// ReleaseJumpAddressHsfRequest 释放JumpAddressHsfRequest
func ReleaseJumpAddressHsfRequest(v *JumpAddressHsfRequest) {
	v.UserNick = ""
	v.AppKey = ""
	v.UserKey = ""
	v.Platform = ""
	poolJumpAddressHsfRequest.Put(v)
}
