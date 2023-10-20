package wdk

import (
	"sync"
)

// BarcodeBo 结构体
type BarcodeBo struct {
	// 条码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 是否主条码:1是主条码，0非主条码
	MainFlag string `json:"main_flag,omitempty" xml:"main_flag,omitempty"`
	// 条码商品规格，6：比如一个条码对应6瓶啤酒
	SpuSpec int64 `json:"spu_spec,omitempty" xml:"spu_spec,omitempty"`
}

var poolBarcodeBo = sync.Pool{
	New: func() any {
		return new(BarcodeBo)
	},
}

// GetBarcodeBo() 从对象池中获取BarcodeBo
func GetBarcodeBo() *BarcodeBo {
	return poolBarcodeBo.Get().(*BarcodeBo)
}

// ReleaseBarcodeBo 释放BarcodeBo
func ReleaseBarcodeBo(v *BarcodeBo) {
	v.Barcode = ""
	v.MainFlag = ""
	v.SpuSpec = 0
	poolBarcodeBo.Put(v)
}
