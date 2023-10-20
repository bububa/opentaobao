package wms

import (
	"sync"
)

// CainiaoStockOutBillPackageinfo 结构体
type CainiaoStockOutBillPackageinfo struct {
	// 包裹里面的商品信息列表
	PackageItemList []CainiaoStockOutBillPackageitemlist `json:"package_item_list,omitempty" xml:"package_item_list>cainiao_stock_out_bill_packageitemlist,omitempty"`
	// 快递公司服务编码
	TmsCode string `json:"tms_code,omitempty" xml:"tms_code,omitempty"`
	// 运单编码
	TmsOrderCode string `json:"tms_order_code,omitempty" xml:"tms_order_code,omitempty"`
	// 包裹号
	PackageCode string `json:"package_code,omitempty" xml:"package_code,omitempty"`
	// 包裹质量，单位：克
	PackageWeight int64 `json:"package_weight,omitempty" xml:"package_weight,omitempty"`
	// 包裹长度，单位：毫米
	PackageLength int64 `json:"package_length,omitempty" xml:"package_length,omitempty"`
	// 包裹宽度,单位：毫米
	PackageWidth int64 `json:"package_width,omitempty" xml:"package_width,omitempty"`
	// 包裹高度，单位：毫米
	PackageHeight int64 `json:"package_height,omitempty" xml:"package_height,omitempty"`
}

var poolCainiaoStockOutBillPackageinfo = sync.Pool{
	New: func() any {
		return new(CainiaoStockOutBillPackageinfo)
	},
}

// GetCainiaoStockOutBillPackageinfo() 从对象池中获取CainiaoStockOutBillPackageinfo
func GetCainiaoStockOutBillPackageinfo() *CainiaoStockOutBillPackageinfo {
	return poolCainiaoStockOutBillPackageinfo.Get().(*CainiaoStockOutBillPackageinfo)
}

// ReleaseCainiaoStockOutBillPackageinfo 释放CainiaoStockOutBillPackageinfo
func ReleaseCainiaoStockOutBillPackageinfo(v *CainiaoStockOutBillPackageinfo) {
	v.PackageItemList = v.PackageItemList[:0]
	v.TmsCode = ""
	v.TmsOrderCode = ""
	v.PackageCode = ""
	v.PackageWeight = 0
	v.PackageLength = 0
	v.PackageWidth = 0
	v.PackageHeight = 0
	poolCainiaoStockOutBillPackageinfo.Put(v)
}
