package alihouse

import (
	"sync"
)

// BaseSupportingDto 结构体
type BaseSupportingDto struct {
	// 外部配套id
	OuterSid string `json:"outer_sid,omitempty" xml:"outer_sid,omitempty"`
	// 外部楼盘、小区、公寓ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 配套类型
	SupportType string `json:"support_type,omitempty" xml:"support_type,omitempty"`
	// 点位编码
	PointCode string `json:"point_code,omitempty" xml:"point_code,omitempty"`
	// 配套名称
	SupportName string `json:"support_name,omitempty" xml:"support_name,omitempty"`
	// 配套地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 配套经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 配套纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 直线距离(单位：米)
	LinearDistance string `json:"linear_distance,omitempty" xml:"linear_distance,omitempty"`
	// 步行距离(单位：米)
	WalkingDistance string `json:"walking_distance,omitempty" xml:"walking_distance,omitempty"`
	// 车行距离(单位：米)
	CarDistance string `json:"car_distance,omitempty" xml:"car_distance,omitempty"`
	// 外部项目店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 1-新房楼盘 2-小区 3-公寓
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 0-无效 1-有效
	IsValid int64 `json:"is_valid,omitempty" xml:"is_valid,omitempty"`
	// 所属省份id
	ProvId int64 `json:"prov_id,omitempty" xml:"prov_id,omitempty"`
	// 所属城市id
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
}

var poolBaseSupportingDto = sync.Pool{
	New: func() any {
		return new(BaseSupportingDto)
	},
}

// GetBaseSupportingDto() 从对象池中获取BaseSupportingDto
func GetBaseSupportingDto() *BaseSupportingDto {
	return poolBaseSupportingDto.Get().(*BaseSupportingDto)
}

// ReleaseBaseSupportingDto 释放BaseSupportingDto
func ReleaseBaseSupportingDto(v *BaseSupportingDto) {
	v.OuterSid = ""
	v.OuterId = ""
	v.SupportType = ""
	v.PointCode = ""
	v.SupportName = ""
	v.Address = ""
	v.Longitude = ""
	v.Latitude = ""
	v.LinearDistance = ""
	v.WalkingDistance = ""
	v.CarDistance = ""
	v.OuterStoreId = ""
	v.Type = 0
	v.IsValid = 0
	v.ProvId = 0
	v.CityId = 0
	poolBaseSupportingDto.Put(v)
}
