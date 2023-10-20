package damai

import (
	"sync"
)

// ThirdPerformPushOpenParam 结构体
type ThirdPerformPushOpenParam struct {
	// 场次结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 场次名称
	PerformName string `json:"perform_name,omitempty" xml:"perform_name,omitempty"`
	// 推送时间
	PushTime string `json:"push_time,omitempty" xml:"push_time,omitempty"`
	// 场次开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 场次id
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 项目id
	ProjectId int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
	// 场馆id
	VenueId int64 `json:"venue_id,omitempty" xml:"venue_id,omitempty"`
	// 入场方式 : 1 = 纸制票入场，2 = 电子票入场，3 = 纸质票电子票均支持
	InWay int64 `json:"in_way,omitempty" xml:"in_way,omitempty"`
}

var poolThirdPerformPushOpenParam = sync.Pool{
	New: func() any {
		return new(ThirdPerformPushOpenParam)
	},
}

// GetThirdPerformPushOpenParam() 从对象池中获取ThirdPerformPushOpenParam
func GetThirdPerformPushOpenParam() *ThirdPerformPushOpenParam {
	return poolThirdPerformPushOpenParam.Get().(*ThirdPerformPushOpenParam)
}

// ReleaseThirdPerformPushOpenParam 释放ThirdPerformPushOpenParam
func ReleaseThirdPerformPushOpenParam(v *ThirdPerformPushOpenParam) {
	v.EndTime = ""
	v.PerformName = ""
	v.PushTime = ""
	v.StartTime = ""
	v.SupplierSecret = ""
	v.PerformId = 0
	v.ProjectId = 0
	v.SystemId = 0
	v.VenueId = 0
	v.InWay = 0
	poolThirdPerformPushOpenParam.Put(v)
}
