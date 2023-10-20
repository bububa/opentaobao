package damai

import (
	"sync"
)

// ThirdVenuePushOpenParam 结构体
type ThirdVenuePushOpenParam struct {
	// 数据推送时间
	PushTime string `json:"push_time,omitempty" xml:"push_time,omitempty"`
	// 商户密钥，监权使用
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 场馆地址
	VenueAddress string `json:"venue_address,omitempty" xml:"venue_address,omitempty"`
	// 场馆名称
	VenueName string `json:"venue_name,omitempty" xml:"venue_name,omitempty"`
	// 商户id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
	// 场馆id
	VenueId int64 `json:"venue_id,omitempty" xml:"venue_id,omitempty"`
}

var poolThirdVenuePushOpenParam = sync.Pool{
	New: func() any {
		return new(ThirdVenuePushOpenParam)
	},
}

// GetThirdVenuePushOpenParam() 从对象池中获取ThirdVenuePushOpenParam
func GetThirdVenuePushOpenParam() *ThirdVenuePushOpenParam {
	return poolThirdVenuePushOpenParam.Get().(*ThirdVenuePushOpenParam)
}

// ReleaseThirdVenuePushOpenParam 释放ThirdVenuePushOpenParam
func ReleaseThirdVenuePushOpenParam(v *ThirdVenuePushOpenParam) {
	v.PushTime = ""
	v.SupplierSecret = ""
	v.VenueAddress = ""
	v.VenueName = ""
	v.SystemId = 0
	v.VenueId = 0
	poolThirdVenuePushOpenParam.Put(v)
}
