package damai

import (
	"sync"
)

// ThirdFloorPushOpenParam 结构体
type ThirdFloorPushOpenParam struct {
	// 楼层名称
	FloorName string `json:"floor_name,omitempty" xml:"floor_name,omitempty"`
	// 推送时间
	PushTime string `json:"push_time,omitempty" xml:"push_time,omitempty"`
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 楼层id
	FloorId int64 `json:"floor_id,omitempty" xml:"floor_id,omitempty"`
	// 场次id
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
	// 场馆id
	VenueId int64 `json:"venue_id,omitempty" xml:"venue_id,omitempty"`
}

var poolThirdFloorPushOpenParam = sync.Pool{
	New: func() any {
		return new(ThirdFloorPushOpenParam)
	},
}

// GetThirdFloorPushOpenParam() 从对象池中获取ThirdFloorPushOpenParam
func GetThirdFloorPushOpenParam() *ThirdFloorPushOpenParam {
	return poolThirdFloorPushOpenParam.Get().(*ThirdFloorPushOpenParam)
}

// ReleaseThirdFloorPushOpenParam 释放ThirdFloorPushOpenParam
func ReleaseThirdFloorPushOpenParam(v *ThirdFloorPushOpenParam) {
	v.FloorName = ""
	v.PushTime = ""
	v.SupplierSecret = ""
	v.FloorId = 0
	v.PerformId = 0
	v.SystemId = 0
	v.VenueId = 0
	poolThirdFloorPushOpenParam.Put(v)
}
