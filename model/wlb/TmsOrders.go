package wlb

import (
	"sync"
)

// TmsOrders 结构体
type TmsOrders struct {
	// 运单信息
	TmsItems []TmsItem `json:"tms_items,omitempty" xml:"tms_items>tms_item,omitempty"`
	// 包裹的包材信息列表
	PackageMaterialList []PackageMaterialList `json:"package_material_list,omitempty" xml:"package_material_list>package_material_list,omitempty"`
	// 运单编码，运单号
	TmsOrderCode string `json:"tms_order_code,omitempty" xml:"tms_order_code,omitempty"`
	// 快递公司服务编码
	TmsCode string `json:"tms_code,omitempty" xml:"tms_code,omitempty"`
	// 包裹号
	PackageCode string `json:"package_code,omitempty" xml:"package_code,omitempty"`
	// 包裹高度，单位：毫米
	PackageHeight int64 `json:"package_height,omitempty" xml:"package_height,omitempty"`
	// 包裹宽度，单位：毫米
	PackageWidth int64 `json:"package_width,omitempty" xml:"package_width,omitempty"`
	// 包裹长度，单位：毫米
	PackageLength int64 `json:"package_length,omitempty" xml:"package_length,omitempty"`
	// 包裹重量，单位：克
	PackageWeight int64 `json:"package_weight,omitempty" xml:"package_weight,omitempty"`
}

var poolTmsOrders = sync.Pool{
	New: func() any {
		return new(TmsOrders)
	},
}

// GetTmsOrders() 从对象池中获取TmsOrders
func GetTmsOrders() *TmsOrders {
	return poolTmsOrders.Get().(*TmsOrders)
}

// ReleaseTmsOrders 释放TmsOrders
func ReleaseTmsOrders(v *TmsOrders) {
	v.TmsItems = v.TmsItems[:0]
	v.PackageMaterialList = v.PackageMaterialList[:0]
	v.TmsOrderCode = ""
	v.TmsCode = ""
	v.PackageCode = ""
	v.PackageHeight = 0
	v.PackageWidth = 0
	v.PackageLength = 0
	v.PackageWeight = 0
	poolTmsOrders.Put(v)
}
