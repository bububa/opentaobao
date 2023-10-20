package tmallnr

import (
	"sync"
)

// Errmsg 结构体
type Errmsg struct {
	// 错误编码
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 订单号
	Key int64 `json:"key,omitempty" xml:"key,omitempty"`
}

var poolErrmsg = sync.Pool{
	New: func() any {
		return new(Errmsg)
	},
}

// GetErrmsg() 从对象池中获取Errmsg
func GetErrmsg() *Errmsg {
	return poolErrmsg.Get().(*Errmsg)
}

// ReleaseErrmsg 释放Errmsg
func ReleaseErrmsg(v *Errmsg) {
	v.Value = ""
	v.Key = 0
	poolErrmsg.Put(v)
}
