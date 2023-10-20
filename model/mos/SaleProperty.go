package mos

import (
	"sync"
)

// SaleProperty 结构体
type SaleProperty struct {
	// 属性Id
	PId string `json:"p_id,omitempty" xml:"p_id,omitempty"`
	// 属性名称
	PName string `json:"p_name,omitempty" xml:"p_name,omitempty"`
	// 属性值Id
	VId string `json:"v_id,omitempty" xml:"v_id,omitempty"`
	// 属性值名称
	VName string `json:"v_name,omitempty" xml:"v_name,omitempty"`
}

var poolSaleProperty = sync.Pool{
	New: func() any {
		return new(SaleProperty)
	},
}

// GetSaleProperty() 从对象池中获取SaleProperty
func GetSaleProperty() *SaleProperty {
	return poolSaleProperty.Get().(*SaleProperty)
}

// ReleaseSaleProperty 释放SaleProperty
func ReleaseSaleProperty(v *SaleProperty) {
	v.PId = ""
	v.PName = ""
	v.VId = ""
	v.VName = ""
	poolSaleProperty.Put(v)
}
