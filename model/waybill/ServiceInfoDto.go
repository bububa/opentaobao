package waybill

import (
	"sync"
)

// ServiceInfoDto 结构体
type ServiceInfoDto struct {
	// 服务属性定义
	ServiceAttributes []ServiceAttributeDto `json:"service_attributes,omitempty" xml:"service_attributes>service_attribute_dto,omitempty"`
	// 服务名称
	ServiceName string `json:"service_name,omitempty" xml:"service_name,omitempty"`
	// 服务编码
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 服务的官方描述，可以用作前端展示
	ServiceDesc string `json:"service_desc,omitempty" xml:"service_desc,omitempty"`
	// 该服务是否为必选服务
	Required bool `json:"required,omitempty" xml:"required,omitempty"`
}

var poolServiceInfoDto = sync.Pool{
	New: func() any {
		return new(ServiceInfoDto)
	},
}

// GetServiceInfoDto() 从对象池中获取ServiceInfoDto
func GetServiceInfoDto() *ServiceInfoDto {
	return poolServiceInfoDto.Get().(*ServiceInfoDto)
}

// ReleaseServiceInfoDto 释放ServiceInfoDto
func ReleaseServiceInfoDto(v *ServiceInfoDto) {
	v.ServiceAttributes = v.ServiceAttributes[:0]
	v.ServiceName = ""
	v.ServiceCode = ""
	v.ServiceDesc = ""
	v.Required = false
	poolServiceInfoDto.Put(v)
}
