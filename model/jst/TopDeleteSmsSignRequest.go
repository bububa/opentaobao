package jst

import (
	"sync"
)

// TopDeleteSmsSignRequest 结构体
type TopDeleteSmsSignRequest struct {
	// 待删除的签名
	SignName string `json:"sign_name,omitempty" xml:"sign_name,omitempty"`
}

var poolTopDeleteSmsSignRequest = sync.Pool{
	New: func() any {
		return new(TopDeleteSmsSignRequest)
	},
}

// GetTopDeleteSmsSignRequest() 从对象池中获取TopDeleteSmsSignRequest
func GetTopDeleteSmsSignRequest() *TopDeleteSmsSignRequest {
	return poolTopDeleteSmsSignRequest.Get().(*TopDeleteSmsSignRequest)
}

// ReleaseTopDeleteSmsSignRequest 释放TopDeleteSmsSignRequest
func ReleaseTopDeleteSmsSignRequest(v *TopDeleteSmsSignRequest) {
	v.SignName = ""
	poolTopDeleteSmsSignRequest.Put(v)
}
