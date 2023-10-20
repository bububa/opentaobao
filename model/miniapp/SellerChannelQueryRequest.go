package miniapp

import (
	"sync"
)

// SellerChannelQueryRequest 结构体
type SellerChannelQueryRequest struct {
	// 渠道id
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
}

var poolSellerChannelQueryRequest = sync.Pool{
	New: func() any {
		return new(SellerChannelQueryRequest)
	},
}

// GetSellerChannelQueryRequest() 从对象池中获取SellerChannelQueryRequest
func GetSellerChannelQueryRequest() *SellerChannelQueryRequest {
	return poolSellerChannelQueryRequest.Get().(*SellerChannelQueryRequest)
}

// ReleaseSellerChannelQueryRequest 释放SellerChannelQueryRequest
func ReleaseSellerChannelQueryRequest(v *SellerChannelQueryRequest) {
	v.Channel = ""
	poolSellerChannelQueryRequest.Put(v)
}
