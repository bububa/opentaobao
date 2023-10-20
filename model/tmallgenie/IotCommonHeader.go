package tmallgenie

import (
	"sync"
)

// IotCommonHeader 结构体
type IotCommonHeader struct {
	// 标准控制协议中的namespace
	Namespace string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// 标准控制协议中的name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 标准控制协议中的messageId
	MessageId string `json:"message_id,omitempty" xml:"message_id,omitempty"`
	// 标准控制协议中的payLoadVersion
	PayLoadVersion int64 `json:"pay_load_version,omitempty" xml:"pay_load_version,omitempty"`
}

var poolIotCommonHeader = sync.Pool{
	New: func() any {
		return new(IotCommonHeader)
	},
}

// GetIotCommonHeader() 从对象池中获取IotCommonHeader
func GetIotCommonHeader() *IotCommonHeader {
	return poolIotCommonHeader.Get().(*IotCommonHeader)
}

// ReleaseIotCommonHeader 释放IotCommonHeader
func ReleaseIotCommonHeader(v *IotCommonHeader) {
	v.Namespace = ""
	v.Name = ""
	v.MessageId = ""
	v.PayLoadVersion = 0
	poolIotCommonHeader.Put(v)
}
