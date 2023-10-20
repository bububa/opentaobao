package wms

import (
	"sync"
)

// Packagematerial 结构体
type Packagematerial struct {
	// 包材的数量
	MaterialQuantity string `json:"material_quantity,omitempty" xml:"material_quantity,omitempty"`
	// 淘宝指定的包材型号
	MaterialType string `json:"material_type,omitempty" xml:"material_type,omitempty"`
}

var poolPackagematerial = sync.Pool{
	New: func() any {
		return new(Packagematerial)
	},
}

// GetPackagematerial() 从对象池中获取Packagematerial
func GetPackagematerial() *Packagematerial {
	return poolPackagematerial.Get().(*Packagematerial)
}

// ReleasePackagematerial 释放Packagematerial
func ReleasePackagematerial(v *Packagematerial) {
	v.MaterialQuantity = ""
	v.MaterialType = ""
	poolPackagematerial.Put(v)
}
