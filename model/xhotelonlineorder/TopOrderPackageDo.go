package xhotelonlineorder

import (
	"sync"
)

// TopOrderPackageDo 结构体
type TopOrderPackageDo struct {
	// 套餐详情
	OrderPackageDetails []TopOrderPackageDetail `json:"order_package_details,omitempty" xml:"order_package_details>top_order_package_detail,omitempty"`
	// 描述
	OrderPackageDesc string `json:"order_package_desc,omitempty" xml:"order_package_desc,omitempty"`
}

var poolTopOrderPackageDo = sync.Pool{
	New: func() any {
		return new(TopOrderPackageDo)
	},
}

// GetTopOrderPackageDo() 从对象池中获取TopOrderPackageDo
func GetTopOrderPackageDo() *TopOrderPackageDo {
	return poolTopOrderPackageDo.Get().(*TopOrderPackageDo)
}

// ReleaseTopOrderPackageDo 释放TopOrderPackageDo
func ReleaseTopOrderPackageDo(v *TopOrderPackageDo) {
	v.OrderPackageDetails = v.OrderPackageDetails[:0]
	v.OrderPackageDesc = ""
	poolTopOrderPackageDo.Put(v)
}
