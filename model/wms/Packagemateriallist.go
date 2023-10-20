package wms

import (
	"sync"
)

// Packagemateriallist 结构体
type Packagemateriallist struct {
	// 包裹包材信息
	PackageMaterial *Packagematerial `json:"package_material,omitempty" xml:"package_material,omitempty"`
}

var poolPackagemateriallist = sync.Pool{
	New: func() any {
		return new(Packagemateriallist)
	},
}

// GetPackagemateriallist() 从对象池中获取Packagemateriallist
func GetPackagemateriallist() *Packagemateriallist {
	return poolPackagemateriallist.Get().(*Packagemateriallist)
}

// ReleasePackagemateriallist 释放Packagemateriallist
func ReleasePackagemateriallist(v *Packagemateriallist) {
	v.PackageMaterial = nil
	poolPackagemateriallist.Put(v)
}
