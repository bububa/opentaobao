package tmallservice

import (
	"sync"
)

// ServiceTypeDto 结构体
type ServiceTypeDto struct {
	// serviceName
	ServiceName string `json:"service_name,omitempty" xml:"service_name,omitempty"`
}

var poolServiceTypeDto = sync.Pool{
	New: func() any {
		return new(ServiceTypeDto)
	},
}

// GetServiceTypeDto() 从对象池中获取ServiceTypeDto
func GetServiceTypeDto() *ServiceTypeDto {
	return poolServiceTypeDto.Get().(*ServiceTypeDto)
}

// ReleaseServiceTypeDto 释放ServiceTypeDto
func ReleaseServiceTypeDto(v *ServiceTypeDto) {
	v.ServiceName = ""
	poolServiceTypeDto.Put(v)
}
