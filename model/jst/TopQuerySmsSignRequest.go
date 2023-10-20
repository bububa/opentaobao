package jst

import (
	"sync"
)

// TopQuerySmsSignRequest 结构体
type TopQuerySmsSignRequest struct {
	// 要查询的签名
	SignName string `json:"sign_name,omitempty" xml:"sign_name,omitempty"`
}

var poolTopQuerySmsSignRequest = sync.Pool{
	New: func() any {
		return new(TopQuerySmsSignRequest)
	},
}

// GetTopQuerySmsSignRequest() 从对象池中获取TopQuerySmsSignRequest
func GetTopQuerySmsSignRequest() *TopQuerySmsSignRequest {
	return poolTopQuerySmsSignRequest.Get().(*TopQuerySmsSignRequest)
}

// ReleaseTopQuerySmsSignRequest 释放TopQuerySmsSignRequest
func ReleaseTopQuerySmsSignRequest(v *TopQuerySmsSignRequest) {
	v.SignName = ""
	poolTopQuerySmsSignRequest.Put(v)
}
