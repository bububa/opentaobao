package wms

import (
	"sync"
)

// Tmsorder 结构体
type Tmsorder struct {
	// 包裹里面的商品信息列表
	TmsItemList []Tmsitemlist `json:"tms_item_list,omitempty" xml:"tms_item_list>tmsitemlist,omitempty"`
	// 包材信息
	PackageMaterialList []Packagemateriallist `json:"package_material_list,omitempty" xml:"package_material_list>packagemateriallist,omitempty"`
	// 包裹号
	PackageCode string `json:"package_code,omitempty" xml:"package_code,omitempty"`
	// 运单编码
	TmsOrderCode string `json:"tms_order_code,omitempty" xml:"tms_order_code,omitempty"`
	// 快递公司服务编码
	TmsCode string `json:"tms_code,omitempty" xml:"tms_code,omitempty"`
	// 包裹高度，单位：毫米
	PackageHeight int64 `json:"package_height,omitempty" xml:"package_height,omitempty"`
	// 包裹宽度，单位：毫米
	PackageWidth int64 `json:"package_width,omitempty" xml:"package_width,omitempty"`
	// 包裹长度，单位：毫米
	PackageLength int64 `json:"package_length,omitempty" xml:"package_length,omitempty"`
	// 包裹重量，单位：克
	PackageWeight int64 `json:"package_weight,omitempty" xml:"package_weight,omitempty"`
}

var poolTmsorder = sync.Pool{
	New: func() any {
		return new(Tmsorder)
	},
}

// GetTmsorder() 从对象池中获取Tmsorder
func GetTmsorder() *Tmsorder {
	return poolTmsorder.Get().(*Tmsorder)
}

// ReleaseTmsorder 释放Tmsorder
func ReleaseTmsorder(v *Tmsorder) {
	v.TmsItemList = v.TmsItemList[:0]
	v.PackageMaterialList = v.PackageMaterialList[:0]
	v.PackageCode = ""
	v.TmsOrderCode = ""
	v.TmsCode = ""
	v.PackageHeight = 0
	v.PackageWidth = 0
	v.PackageLength = 0
	v.PackageWeight = 0
	poolTmsorder.Put(v)
}
