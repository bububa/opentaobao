package jst

import (
	"sync"
)

// SmsSignNameRequest 结构体
type SmsSignNameRequest struct {
	// 短信签名
	SignName string `json:"sign_name,omitempty" xml:"sign_name,omitempty"`
	// 描述信息
	Description string `json:"description,omitempty" xml:"description,omitempty"`
}

var poolSmsSignNameRequest = sync.Pool{
	New: func() any {
		return new(SmsSignNameRequest)
	},
}

// GetSmsSignNameRequest() 从对象池中获取SmsSignNameRequest
func GetSmsSignNameRequest() *SmsSignNameRequest {
	return poolSmsSignNameRequest.Get().(*SmsSignNameRequest)
}

// ReleaseSmsSignNameRequest 释放SmsSignNameRequest
func ReleaseSmsSignNameRequest(v *SmsSignNameRequest) {
	v.SignName = ""
	v.Description = ""
	poolSmsSignNameRequest.Put(v)
}
