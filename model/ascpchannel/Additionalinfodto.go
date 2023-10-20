package ascpchannel

import (
	"sync"
)

// Additionalinfodto 结构体
type Additionalinfodto struct {
	// 未税价格 (精度 6位) 采购&amp;退供 经销必填
	PriceNoTax string `json:"price_no_tax,omitempty" xml:"price_no_tax,omitempty"`
}

var poolAdditionalinfodto = sync.Pool{
	New: func() any {
		return new(Additionalinfodto)
	},
}

// GetAdditionalinfodto() 从对象池中获取Additionalinfodto
func GetAdditionalinfodto() *Additionalinfodto {
	return poolAdditionalinfodto.Get().(*Additionalinfodto)
}

// ReleaseAdditionalinfodto 释放Additionalinfodto
func ReleaseAdditionalinfodto(v *Additionalinfodto) {
	v.PriceNoTax = ""
	poolAdditionalinfodto.Put(v)
}
