package happytrip

import (
	"sync"
)

// ServiceStatusModel 结构体
type ServiceStatusModel struct {
	// 供应商服务在各地区的支持状态
	Cities []CityServiceStatus `json:"cities,omitempty" xml:"cities>city_service_status,omitempty"`
}

var poolServiceStatusModel = sync.Pool{
	New: func() any {
		return new(ServiceStatusModel)
	},
}

// GetServiceStatusModel() 从对象池中获取ServiceStatusModel
func GetServiceStatusModel() *ServiceStatusModel {
	return poolServiceStatusModel.Get().(*ServiceStatusModel)
}

// ReleaseServiceStatusModel 释放ServiceStatusModel
func ReleaseServiceStatusModel(v *ServiceStatusModel) {
	v.Cities = v.Cities[:0]
	poolServiceStatusModel.Put(v)
}
