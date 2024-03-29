package damai

import (
	"sync"
)

// ThirdStandPushOpenParam 结构体
type ThirdStandPushOpenParam struct {
	// 推送时间
	PushTime string `json:"push_time,omitempty" xml:"push_time,omitempty"`
	// 看台名称
	StandName string `json:"stand_name,omitempty" xml:"stand_name,omitempty"`
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 场次id
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 看台id
	StandId int64 `json:"stand_id,omitempty" xml:"stand_id,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
	// 场馆id
	VenueId int64 `json:"venue_id,omitempty" xml:"venue_id,omitempty"`
}

var poolThirdStandPushOpenParam = sync.Pool{
	New: func() any {
		return new(ThirdStandPushOpenParam)
	},
}

// GetThirdStandPushOpenParam() 从对象池中获取ThirdStandPushOpenParam
func GetThirdStandPushOpenParam() *ThirdStandPushOpenParam {
	return poolThirdStandPushOpenParam.Get().(*ThirdStandPushOpenParam)
}

// ReleaseThirdStandPushOpenParam 释放ThirdStandPushOpenParam
func ReleaseThirdStandPushOpenParam(v *ThirdStandPushOpenParam) {
	v.PushTime = ""
	v.StandName = ""
	v.SupplierSecret = ""
	v.PerformId = 0
	v.StandId = 0
	v.SystemId = 0
	v.VenueId = 0
	poolThirdStandPushOpenParam.Put(v)
}
