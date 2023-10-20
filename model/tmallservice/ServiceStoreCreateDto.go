package tmallservice

import (
	"sync"
)

// ServiceStoreCreateDto 结构体
type ServiceStoreCreateDto struct {
	// 秘钥--内嵌核销页面使用
	Secret string `json:"secret,omitempty" xml:"secret,omitempty"`
	// 网点id
	ServiceStoreId int64 `json:"service_store_id,omitempty" xml:"service_store_id,omitempty"`
}

var poolServiceStoreCreateDto = sync.Pool{
	New: func() any {
		return new(ServiceStoreCreateDto)
	},
}

// GetServiceStoreCreateDto() 从对象池中获取ServiceStoreCreateDto
func GetServiceStoreCreateDto() *ServiceStoreCreateDto {
	return poolServiceStoreCreateDto.Get().(*ServiceStoreCreateDto)
}

// ReleaseServiceStoreCreateDto 释放ServiceStoreCreateDto
func ReleaseServiceStoreCreateDto(v *ServiceStoreCreateDto) {
	v.Secret = ""
	v.ServiceStoreId = 0
	poolServiceStoreCreateDto.Put(v)
}
