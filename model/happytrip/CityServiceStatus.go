package happytrip

import (
	"sync"
)

// CityServiceStatus 结构体
type CityServiceStatus struct {
	// 支持的车型代码列表
	SupportServiceLevels []string `json:"support_service_levels,omitempty" xml:"support_service_levels>string,omitempty"`
	// 供应商的城市id
	CityId string `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 城市名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolCityServiceStatus = sync.Pool{
	New: func() any {
		return new(CityServiceStatus)
	},
}

// GetCityServiceStatus() 从对象池中获取CityServiceStatus
func GetCityServiceStatus() *CityServiceStatus {
	return poolCityServiceStatus.Get().(*CityServiceStatus)
}

// ReleaseCityServiceStatus 释放CityServiceStatus
func ReleaseCityServiceStatus(v *CityServiceStatus) {
	v.SupportServiceLevels = v.SupportServiceLevels[:0]
	v.CityId = ""
	v.Name = ""
	poolCityServiceStatus.Put(v)
}
