package damai

import (
	"sync"
)

// ThirdProjectPushOpenParam 结构体
type ThirdProjectPushOpenParam struct {
	// 图片url
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 项目名称
	ProjectName string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// 推送时间
	PushTime string `json:"push_time,omitempty" xml:"push_time,omitempty"`
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 城市id
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 项目id
	ProjectId int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
	// 场馆id
	VenueId int64 `json:"venue_id,omitempty" xml:"venue_id,omitempty"`
}

var poolThirdProjectPushOpenParam = sync.Pool{
	New: func() any {
		return new(ThirdProjectPushOpenParam)
	},
}

// GetThirdProjectPushOpenParam() 从对象池中获取ThirdProjectPushOpenParam
func GetThirdProjectPushOpenParam() *ThirdProjectPushOpenParam {
	return poolThirdProjectPushOpenParam.Get().(*ThirdProjectPushOpenParam)
}

// ReleaseThirdProjectPushOpenParam 释放ThirdProjectPushOpenParam
func ReleaseThirdProjectPushOpenParam(v *ThirdProjectPushOpenParam) {
	v.PicUrl = ""
	v.ProjectName = ""
	v.PushTime = ""
	v.SupplierSecret = ""
	v.CityId = 0
	v.ProjectId = 0
	v.SystemId = 0
	v.VenueId = 0
	poolThirdProjectPushOpenParam.Put(v)
}
