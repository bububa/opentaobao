package wms

import (
	"sync"
)

// CainiaoStockOutBillPackageitemlist 结构体
type CainiaoStockOutBillPackageitemlist struct {
	// 1
	PackageItem *CainiaoStockOutBillPackageitem `json:"package_item,omitempty" xml:"package_item,omitempty"`
}

var poolCainiaoStockOutBillPackageitemlist = sync.Pool{
	New: func() any {
		return new(CainiaoStockOutBillPackageitemlist)
	},
}

// GetCainiaoStockOutBillPackageitemlist() 从对象池中获取CainiaoStockOutBillPackageitemlist
func GetCainiaoStockOutBillPackageitemlist() *CainiaoStockOutBillPackageitemlist {
	return poolCainiaoStockOutBillPackageitemlist.Get().(*CainiaoStockOutBillPackageitemlist)
}

// ReleaseCainiaoStockOutBillPackageitemlist 释放CainiaoStockOutBillPackageitemlist
func ReleaseCainiaoStockOutBillPackageitemlist(v *CainiaoStockOutBillPackageitemlist) {
	v.PackageItem = nil
	poolCainiaoStockOutBillPackageitemlist.Put(v)
}
