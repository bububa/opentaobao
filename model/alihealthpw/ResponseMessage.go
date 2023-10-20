package alihealthpw

import (
	"sync"
)

// ResponseMessage 结构体
type ResponseMessage struct {
	// 状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 数据
	Data *ApplyDetailResp `json:"data,omitempty" xml:"data,omitempty"`
}

var poolResponseMessage = sync.Pool{
	New: func() any {
		return new(ResponseMessage)
	},
}

// GetResponseMessage() 从对象池中获取ResponseMessage
func GetResponseMessage() *ResponseMessage {
	return poolResponseMessage.Get().(*ResponseMessage)
}

// ReleaseResponseMessage 释放ResponseMessage
func ReleaseResponseMessage(v *ResponseMessage) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	poolResponseMessage.Put(v)
}
