package icbulogistics

import (
	"sync"
)

// PackageList 结构体
type PackageList struct {
	// 数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 长
	Length string `json:"length,omitempty" xml:"length,omitempty"`
	// 宽
	Width string `json:"width,omitempty" xml:"width,omitempty"`
	// 重量
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 包装类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 高
	Height string `json:"height,omitempty" xml:"height,omitempty"`
}

var poolPackageList = sync.Pool{
	New: func() any {
		return new(PackageList)
	},
}

// GetPackageList() 从对象池中获取PackageList
func GetPackageList() *PackageList {
	return poolPackageList.Get().(*PackageList)
}

// ReleasePackageList 释放PackageList
func ReleasePackageList(v *PackageList) {
	v.Quantity = ""
	v.Length = ""
	v.Width = ""
	v.Weight = ""
	v.Type = ""
	v.Height = ""
	poolPackageList.Put(v)
}
