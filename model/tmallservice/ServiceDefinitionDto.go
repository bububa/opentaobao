package tmallservice

import (
	"sync"
)

// ServiceDefinitionDto 结构体
type ServiceDefinitionDto struct {
	// 服务编码
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 服务名称
	DisplayName string `json:"display_name,omitempty" xml:"display_name,omitempty"`
	// 业务编码
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
}

var poolServiceDefinitionDto = sync.Pool{
	New: func() any {
		return new(ServiceDefinitionDto)
	},
}

// GetServiceDefinitionDto() 从对象池中获取ServiceDefinitionDto
func GetServiceDefinitionDto() *ServiceDefinitionDto {
	return poolServiceDefinitionDto.Get().(*ServiceDefinitionDto)
}

// ReleaseServiceDefinitionDto 释放ServiceDefinitionDto
func ReleaseServiceDefinitionDto(v *ServiceDefinitionDto) {
	v.ServiceCode = ""
	v.DisplayName = ""
	v.BizCode = ""
	poolServiceDefinitionDto.Put(v)
}
