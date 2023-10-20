package wms

import (
	"sync"
)

// CainiaoStockOutBillPackageinfolist 结构体
type CainiaoStockOutBillPackageinfolist struct {
	// 包裹信息
	PackageInfo *CainiaoStockOutBillPackageinfo `json:"package_info,omitempty" xml:"package_info,omitempty"`
}

var poolCainiaoStockOutBillPackageinfolist = sync.Pool{
	New: func() any {
		return new(CainiaoStockOutBillPackageinfolist)
	},
}

// GetCainiaoStockOutBillPackageinfolist() 从对象池中获取CainiaoStockOutBillPackageinfolist
func GetCainiaoStockOutBillPackageinfolist() *CainiaoStockOutBillPackageinfolist {
	return poolCainiaoStockOutBillPackageinfolist.Get().(*CainiaoStockOutBillPackageinfolist)
}

// ReleaseCainiaoStockOutBillPackageinfolist 释放CainiaoStockOutBillPackageinfolist
func ReleaseCainiaoStockOutBillPackageinfolist(v *CainiaoStockOutBillPackageinfolist) {
	v.PackageInfo = nil
	poolCainiaoStockOutBillPackageinfolist.Put(v)
}
