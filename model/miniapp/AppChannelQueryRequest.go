package miniapp

import (
	"sync"
)

// AppChannelQueryRequest 结构体
type AppChannelQueryRequest struct {
	// 渠道id
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 小程序id
	MiniappId int64 `json:"miniapp_id,omitempty" xml:"miniapp_id,omitempty"`
}

var poolAppChannelQueryRequest = sync.Pool{
	New: func() any {
		return new(AppChannelQueryRequest)
	},
}

// GetAppChannelQueryRequest() 从对象池中获取AppChannelQueryRequest
func GetAppChannelQueryRequest() *AppChannelQueryRequest {
	return poolAppChannelQueryRequest.Get().(*AppChannelQueryRequest)
}

// ReleaseAppChannelQueryRequest 释放AppChannelQueryRequest
func ReleaseAppChannelQueryRequest(v *AppChannelQueryRequest) {
	v.Channel = ""
	v.MiniappId = 0
	poolAppChannelQueryRequest.Put(v)
}
