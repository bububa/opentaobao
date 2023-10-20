package koubeimall

import (
	"sync"
)

// DistrictInfo 结构体
type DistrictInfo struct {
	// 区域编码
	DistrictCode string `json:"district_code,omitempty" xml:"district_code,omitempty"`
	// 地址信息
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 区域名称
	DistrictName string `json:"district_name,omitempty" xml:"district_name,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 城市编码
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 省份编码
	ProvinceCode string `json:"province_code,omitempty" xml:"province_code,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 省份名称
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
}

var poolDistrictInfo = sync.Pool{
	New: func() any {
		return new(DistrictInfo)
	},
}

// GetDistrictInfo() 从对象池中获取DistrictInfo
func GetDistrictInfo() *DistrictInfo {
	return poolDistrictInfo.Get().(*DistrictInfo)
}

// ReleaseDistrictInfo 释放DistrictInfo
func ReleaseDistrictInfo(v *DistrictInfo) {
	v.DistrictCode = ""
	v.Address = ""
	v.DistrictName = ""
	v.CityName = ""
	v.CityCode = ""
	v.ProvinceCode = ""
	v.Latitude = ""
	v.ProvinceName = ""
	v.Longitude = ""
	poolDistrictInfo.Put(v)
}
