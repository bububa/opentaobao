package wdkitem

import (
	"sync"
)

// Barcodebolist 结构体
type Barcodebolist struct {
	// barcode
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 规格
	SpuSpec string `json:"spu_spec,omitempty" xml:"spu_spec,omitempty"`
}

var poolBarcodebolist = sync.Pool{
	New: func() any {
		return new(Barcodebolist)
	},
}

// GetBarcodebolist() 从对象池中获取Barcodebolist
func GetBarcodebolist() *Barcodebolist {
	return poolBarcodebolist.Get().(*Barcodebolist)
}

// ReleaseBarcodebolist 释放Barcodebolist
func ReleaseBarcodebolist(v *Barcodebolist) {
	v.Barcode = ""
	v.SpuSpec = ""
	poolBarcodebolist.Put(v)
}
