package travel

import (
	"sync"
)

// ItemScenic 结构体
type ItemScenic struct {
	// 关联的套餐id
	RelatedPackageId string `json:"related_package_id,omitempty" xml:"related_package_id,omitempty"`
	// 必填，门票类型
	TicketType string `json:"ticket_type,omitempty" xml:"ticket_type,omitempty"`
	// POI来源，AMAP/GOOGLE。境内为高德（AMAP） 境外为GOOGLE
	PoiResource string `json:"poi_resource,omitempty" xml:"poi_resource,omitempty"`
	// 景点的经纬度信息，经度在前，纬度在后，英文逗号分隔 最多支持6位小数，如120.111222,30.111222
	Poi string `json:"poi,omitempty" xml:"poi,omitempty"`
	// 必填，景点名称
	CnName string `json:"cn_name,omitempty" xml:"cn_name,omitempty"`
	// 必填，景点所在城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
}

var poolItemScenic = sync.Pool{
	New: func() any {
		return new(ItemScenic)
	},
}

// GetItemScenic() 从对象池中获取ItemScenic
func GetItemScenic() *ItemScenic {
	return poolItemScenic.Get().(*ItemScenic)
}

// ReleaseItemScenic 释放ItemScenic
func ReleaseItemScenic(v *ItemScenic) {
	v.RelatedPackageId = ""
	v.TicketType = ""
	v.PoiResource = ""
	v.Poi = ""
	v.CnName = ""
	v.City = ""
	poolItemScenic.Put(v)
}
