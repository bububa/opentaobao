package alilabs

import (
	"sync"
)

// IotCommonHeader 结构体
type IotCommonHeader struct {
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// messageId，区分请求使用
	MessageId string `json:"message_id,omitempty" xml:"message_id,omitempty"`
	// namespace
	Namespace string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// payLoadVersion
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
	v.Name = ""
	v.MessageId = ""
	v.Namespace = ""
	v.PayLoadVersion = 0
	poolIotCommonHeader.Put(v)
}
