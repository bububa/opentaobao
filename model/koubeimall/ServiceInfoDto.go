package koubeimall

import (
	"sync"
)

// ServiceInfoDto 结构体
type ServiceInfoDto struct {
	// 门店服务tag
	ServiceTagList []string `json:"service_tag_list,omitempty" xml:"service_tag_list>string,omitempty"`
	// 服务说明
	ServiceDesc string `json:"service_desc,omitempty" xml:"service_desc,omitempty"`
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
	v.ServiceTagList = v.ServiceTagList[:0]
	v.ServiceDesc = ""
	poolServiceInfoDto.Put(v)
}
