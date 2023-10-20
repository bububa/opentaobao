package qimen

import (
	"sync"
)

// PackageMaterial 结构体
type PackageMaterial struct {
	// 包材型号
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 包材的数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolPackageMaterial = sync.Pool{
	New: func() any {
		return new(PackageMaterial)
	},
}

// GetPackageMaterial() 从对象池中获取PackageMaterial
func GetPackageMaterial() *PackageMaterial {
	return poolPackageMaterial.Get().(*PackageMaterial)
}

// ReleasePackageMaterial 释放PackageMaterial
func ReleasePackageMaterial(v *PackageMaterial) {
	v.Type = ""
	v.Quantity = 0
	poolPackageMaterial.Put(v)
}
