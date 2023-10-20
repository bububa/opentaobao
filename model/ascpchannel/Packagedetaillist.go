package ascpchannel

import (
	"sync"
)

// Packagedetaillist 结构体
type Packagedetaillist struct {
	// 包件体积(m3)
	PackageVolume string `json:"package_volume,omitempty" xml:"package_volume,omitempty"`
	// 包件重量(kg)
	PackageWeight string `json:"package_weight,omitempty" xml:"package_weight,omitempty"`
	// 包件名称
	PackageName string `json:"package_name,omitempty" xml:"package_name,omitempty"`
	// 扩展字段JSON串
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
}

var poolPackagedetaillist = sync.Pool{
	New: func() any {
		return new(Packagedetaillist)
	},
}

// GetPackagedetaillist() 从对象池中获取Packagedetaillist
func GetPackagedetaillist() *Packagedetaillist {
	return poolPackagedetaillist.Get().(*Packagedetaillist)
}

// ReleasePackagedetaillist 释放Packagedetaillist
func ReleasePackagedetaillist(v *Packagedetaillist) {
	v.PackageVolume = ""
	v.PackageWeight = ""
	v.PackageName = ""
	v.Feature = ""
	poolPackagedetaillist.Put(v)
}
