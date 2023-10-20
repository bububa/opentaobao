package newretail

import (
	"sync"
)

// ApAddressInfo 结构体
type ApAddressInfo struct {
	// ap mac
	Mac string `json:"mac,omitempty" xml:"mac,omitempty"`
	// ap的名称
	ApName string `json:"ap_name,omitempty" xml:"ap_name,omitempty"`
	// ap所在组名
	ApGroup string `json:"ap_group,omitempty" xml:"ap_group,omitempty"`
	// 国家
	ApNationName string `json:"ap_nation_name,omitempty" xml:"ap_nation_name,omitempty"`
	// 省份
	ApProvinceName string `json:"ap_province_name,omitempty" xml:"ap_province_name,omitempty"`
	// 城市
	ApCityName string `json:"ap_city_name,omitempty" xml:"ap_city_name,omitempty"`
	// 区域
	ApAreaName string `json:"ap_area_name,omitempty" xml:"ap_area_name,omitempty"`
	// 园区/门店
	ApCampusName string `json:"ap_campus_name,omitempty" xml:"ap_campus_name,omitempty"`
	// 楼栋
	ApBuildingName string `json:"ap_building_name,omitempty" xml:"ap_building_name,omitempty"`
	// 楼层
	ApFloor string `json:"ap_floor,omitempty" xml:"ap_floor,omitempty"`
	// 空间单元名称
	ApUnitName string `json:"ap_unit_name,omitempty" xml:"ap_unit_name,omitempty"`
	// 方位
	Direction string `json:"direction,omitempty" xml:"direction,omitempty"`
	// 经度
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
	// 纬度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 语言
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// 空间单元id
	ApUnitId int64 `json:"ap_unit_id,omitempty" xml:"ap_unit_id,omitempty"`
}

var poolApAddressInfo = sync.Pool{
	New: func() any {
		return new(ApAddressInfo)
	},
}

// GetApAddressInfo() 从对象池中获取ApAddressInfo
func GetApAddressInfo() *ApAddressInfo {
	return poolApAddressInfo.Get().(*ApAddressInfo)
}

// ReleaseApAddressInfo 释放ApAddressInfo
func ReleaseApAddressInfo(v *ApAddressInfo) {
	v.Mac = ""
	v.ApName = ""
	v.ApGroup = ""
	v.ApNationName = ""
	v.ApProvinceName = ""
	v.ApCityName = ""
	v.ApAreaName = ""
	v.ApCampusName = ""
	v.ApBuildingName = ""
	v.ApFloor = ""
	v.ApUnitName = ""
	v.Direction = ""
	v.Lng = ""
	v.Lat = ""
	v.Language = ""
	v.ApUnitId = 0
	poolApAddressInfo.Put(v)
}
