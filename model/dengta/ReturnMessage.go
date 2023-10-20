package dengta

import (
	"sync"
)

// ReturnMessage 结构体
type ReturnMessage struct {
	// 错误信息
	Dev string `json:"dev,omitempty" xml:"dev,omitempty"`
	// 错误信息
	Cust string `json:"cust,omitempty" xml:"cust,omitempty"`
}

var poolReturnMessage = sync.Pool{
	New: func() any {
		return new(ReturnMessage)
	},
}

// GetReturnMessage() 从对象池中获取ReturnMessage
func GetReturnMessage() *ReturnMessage {
	return poolReturnMessage.Get().(*ReturnMessage)
}

// ReleaseReturnMessage 释放ReturnMessage
func ReleaseReturnMessage(v *ReturnMessage) {
	v.Dev = ""
	v.Cust = ""
	poolReturnMessage.Put(v)
}
