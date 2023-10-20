package tmallservice

import (
	"sync"
)

// ServiceCapacityAdjustReqDto 结构体
type ServiceCapacityAdjustReqDto struct {
	// 具体的调整值参数
	DayQuantityList []DayQuantity `json:"day_quantity_list,omitempty" xml:"day_quantity_list>day_quantity,omitempty"`
	// 服务提供者
	ServiceProviderDto *ServiceProviderDto `json:"service_provider_dto,omitempty" xml:"service_provider_dto,omitempty"`
}

var poolServiceCapacityAdjustReqDto = sync.Pool{
	New: func() any {
		return new(ServiceCapacityAdjustReqDto)
	},
}

// GetServiceCapacityAdjustReqDto() 从对象池中获取ServiceCapacityAdjustReqDto
func GetServiceCapacityAdjustReqDto() *ServiceCapacityAdjustReqDto {
	return poolServiceCapacityAdjustReqDto.Get().(*ServiceCapacityAdjustReqDto)
}

// ReleaseServiceCapacityAdjustReqDto 释放ServiceCapacityAdjustReqDto
func ReleaseServiceCapacityAdjustReqDto(v *ServiceCapacityAdjustReqDto) {
	v.DayQuantityList = v.DayQuantityList[:0]
	v.ServiceProviderDto = nil
	poolServiceCapacityAdjustReqDto.Put(v)
}
