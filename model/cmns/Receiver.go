package cmns

import (
	"sync"
)

// Receiver 结构体
type Receiver struct {
	// deviceToken值，最多1000个
	Data []string `json:"data,omitempty" xml:"data>string,omitempty"`
	// deviceToken值，最多1000个
	DataList []string `json:"data_list,omitempty" xml:"data_list>string,omitempty"`
	// 只能填写deviceToken
	Type string `json:"type,omitempty" xml:"type,omitempty"`
}

var poolReceiver = sync.Pool{
	New: func() any {
		return new(Receiver)
	},
}

// GetReceiver() 从对象池中获取Receiver
func GetReceiver() *Receiver {
	return poolReceiver.Get().(*Receiver)
}

// ReleaseReceiver 释放Receiver
func ReleaseReceiver(v *Receiver) {
	v.Data = v.Data[:0]
	v.DataList = v.DataList[:0]
	v.Type = ""
	poolReceiver.Put(v)
}
