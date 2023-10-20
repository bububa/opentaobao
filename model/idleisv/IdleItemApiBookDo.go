package idleisv

import (
	"sync"
)

// IdleItemApiBookDo 结构体
type IdleItemApiBookDo struct {
	// 图书ISBN码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 图书ISBN码对应的书名等信息
	BarcodeName string `json:"barcode_name,omitempty" xml:"barcode_name,omitempty"`
}

var poolIdleItemApiBookDo = sync.Pool{
	New: func() any {
		return new(IdleItemApiBookDo)
	},
}

// GetIdleItemApiBookDo() 从对象池中获取IdleItemApiBookDo
func GetIdleItemApiBookDo() *IdleItemApiBookDo {
	return poolIdleItemApiBookDo.Get().(*IdleItemApiBookDo)
}

// ReleaseIdleItemApiBookDo 释放IdleItemApiBookDo
func ReleaseIdleItemApiBookDo(v *IdleItemApiBookDo) {
	v.Barcode = ""
	v.BarcodeName = ""
	poolIdleItemApiBookDo.Put(v)
}
