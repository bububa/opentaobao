package cntms

import (
	"sync"
)

// CnTmsLogisticsOrderItemPackageInfo 结构体
type CnTmsLogisticsOrderItemPackageInfo struct {
	// 发货商品信息，最大50条记录
	Items []CnTmsLogisticsOrderItem `json:"items,omitempty" xml:"items>cn_tms_logistics_order_item,omitempty"`
	// 运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 包裹重量（克）
	PackageWeight string `json:"package_weight,omitempty" xml:"package_weight,omitempty"`
	// 包裹长度（厘米）
	PackageLength string `json:"package_length,omitempty" xml:"package_length,omitempty"`
	// 包裹宽度（厘米）
	PackageWidth string `json:"package_width,omitempty" xml:"package_width,omitempty"`
	// 包裹高度（厘米）
	PackageHeight string `json:"package_height,omitempty" xml:"package_height,omitempty"`
	// 包裹体积（立方厘米）
	PackageVolume string `json:"package_volume,omitempty" xml:"package_volume,omitempty"`
}

var poolCnTmsLogisticsOrderItemPackageInfo = sync.Pool{
	New: func() any {
		return new(CnTmsLogisticsOrderItemPackageInfo)
	},
}

// GetCnTmsLogisticsOrderItemPackageInfo() 从对象池中获取CnTmsLogisticsOrderItemPackageInfo
func GetCnTmsLogisticsOrderItemPackageInfo() *CnTmsLogisticsOrderItemPackageInfo {
	return poolCnTmsLogisticsOrderItemPackageInfo.Get().(*CnTmsLogisticsOrderItemPackageInfo)
}

// ReleaseCnTmsLogisticsOrderItemPackageInfo 释放CnTmsLogisticsOrderItemPackageInfo
func ReleaseCnTmsLogisticsOrderItemPackageInfo(v *CnTmsLogisticsOrderItemPackageInfo) {
	v.Items = v.Items[:0]
	v.MailNo = ""
	v.PackageWeight = ""
	v.PackageLength = ""
	v.PackageWidth = ""
	v.PackageHeight = ""
	v.PackageVolume = ""
	poolCnTmsLogisticsOrderItemPackageInfo.Put(v)
}
