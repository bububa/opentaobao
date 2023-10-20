package ascpchannel

import (
	"sync"
)

// Attribute 结构体
type Attribute struct {
	// 供应商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// FULLAMOUNT:全量模式；INCREMENT:增量模式
	InvOperateMode string `json:"inv_operate_mode,omitempty" xml:"inv_operate_mode,omitempty"`
}

var poolAttribute = sync.Pool{
	New: func() any {
		return new(Attribute)
	},
}

// GetAttribute() 从对象池中获取Attribute
func GetAttribute() *Attribute {
	return poolAttribute.Get().(*Attribute)
}

// ReleaseAttribute 释放Attribute
func ReleaseAttribute(v *Attribute) {
	v.SupplierId = ""
	v.InvOperateMode = ""
	poolAttribute.Put(v)
}
