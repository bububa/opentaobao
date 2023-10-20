package tmallservice

import (
	"sync"
)

// ServiceTypeDtOs 结构体
type ServiceTypeDtOs struct {
	// serviceTypeDTOs
	ServiceTypeDTOs []ServiceTypeDto `json:"service_type_d_t_os,omitempty" xml:"service_type_d_t_os>service_type_dto,omitempty"`
}

var poolServiceTypeDtOs = sync.Pool{
	New: func() any {
		return new(ServiceTypeDtOs)
	},
}

// GetServiceTypeDtOs() 从对象池中获取ServiceTypeDtOs
func GetServiceTypeDtOs() *ServiceTypeDtOs {
	return poolServiceTypeDtOs.Get().(*ServiceTypeDtOs)
}

// ReleaseServiceTypeDtOs 释放ServiceTypeDtOs
func ReleaseServiceTypeDtOs(v *ServiceTypeDtOs) {
	v.ServiceTypeDTOs = v.ServiceTypeDTOs[:0]
	poolServiceTypeDtOs.Put(v)
}
