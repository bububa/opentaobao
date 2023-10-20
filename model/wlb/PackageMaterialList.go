package wlb

import (
	"sync"
)

// PackageMaterialList 结构体
type PackageMaterialList struct {
	// 包材
	MaterialType string `json:"material_type,omitempty" xml:"material_type,omitempty"`
	// 数量
	MaterialQuantity int64 `json:"material_quantity,omitempty" xml:"material_quantity,omitempty"`
}

var poolPackageMaterialList = sync.Pool{
	New: func() any {
		return new(PackageMaterialList)
	},
}

// GetPackageMaterialList() 从对象池中获取PackageMaterialList
func GetPackageMaterialList() *PackageMaterialList {
	return poolPackageMaterialList.Get().(*PackageMaterialList)
}

// ReleasePackageMaterialList 释放PackageMaterialList
func ReleasePackageMaterialList(v *PackageMaterialList) {
	v.MaterialType = ""
	v.MaterialQuantity = 0
	poolPackageMaterialList.Put(v)
}
