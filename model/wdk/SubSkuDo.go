package wdk

import (
	"sync"
)

// SubSkuDo 结构体
type SubSkuDo struct {
	// 子商品编码（需要先有子商品）
	SubSkuCode string `json:"sub_sku_code,omitempty" xml:"sub_sku_code,omitempty"`
	// 子商品数量
	SubSkuNum int64 `json:"sub_sku_num,omitempty" xml:"sub_sku_num,omitempty"`
}

var poolSubSkuDo = sync.Pool{
	New: func() any {
		return new(SubSkuDo)
	},
}

// GetSubSkuDo() 从对象池中获取SubSkuDo
func GetSubSkuDo() *SubSkuDo {
	return poolSubSkuDo.Get().(*SubSkuDo)
}

// ReleaseSubSkuDo 释放SubSkuDo
func ReleaseSubSkuDo(v *SubSkuDo) {
	v.SubSkuCode = ""
	v.SubSkuNum = 0
	poolSubSkuDo.Put(v)
}
