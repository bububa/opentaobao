package eticket

import (
	"sync"
)

// SendMaCallbackResp 结构体
type SendMaCallbackResp struct {
	// 回复业务KV
	AttributeMap string `json:"attribute_map,omitempty" xml:"attribute_map,omitempty"`
}

var poolSendMaCallbackResp = sync.Pool{
	New: func() any {
		return new(SendMaCallbackResp)
	},
}

// GetSendMaCallbackResp() 从对象池中获取SendMaCallbackResp
func GetSendMaCallbackResp() *SendMaCallbackResp {
	return poolSendMaCallbackResp.Get().(*SendMaCallbackResp)
}

// ReleaseSendMaCallbackResp 释放SendMaCallbackResp
func ReleaseSendMaCallbackResp(v *SendMaCallbackResp) {
	v.AttributeMap = ""
	poolSendMaCallbackResp.Put(v)
}
