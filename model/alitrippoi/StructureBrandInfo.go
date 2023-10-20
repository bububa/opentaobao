package alitrippoi

import (
	"sync"
)

// StructureBrandInfo 结构体
type StructureBrandInfo struct {
	// 品牌名
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 品牌英文名
	BrandEnName string `json:"brand_en_name,omitempty" xml:"brand_en_name,omitempty"`
	// 品牌Logo
	BrandLogo string `json:"brand_logo,omitempty" xml:"brand_logo,omitempty"`
}

var poolStructureBrandInfo = sync.Pool{
	New: func() any {
		return new(StructureBrandInfo)
	},
}

// GetStructureBrandInfo() 从对象池中获取StructureBrandInfo
func GetStructureBrandInfo() *StructureBrandInfo {
	return poolStructureBrandInfo.Get().(*StructureBrandInfo)
}

// ReleaseStructureBrandInfo 释放StructureBrandInfo
func ReleaseStructureBrandInfo(v *StructureBrandInfo) {
	v.BrandName = ""
	v.BrandEnName = ""
	v.BrandLogo = ""
	poolStructureBrandInfo.Put(v)
}
