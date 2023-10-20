package drugtrace

import (
	"sync"
)

// BaseInfosDto 结构体
type BaseInfosDto struct {
	// 药品基础信息
	BaseInfoList []BaseInfoDto `json:"base_info_list,omitempty" xml:"base_info_list>base_info_dto,omitempty"`
}

var poolBaseInfosDto = sync.Pool{
	New: func() any {
		return new(BaseInfosDto)
	},
}

// GetBaseInfosDto() 从对象池中获取BaseInfosDto
func GetBaseInfosDto() *BaseInfosDto {
	return poolBaseInfosDto.Get().(*BaseInfosDto)
}

// ReleaseBaseInfosDto 释放BaseInfosDto
func ReleaseBaseInfosDto(v *BaseInfosDto) {
	v.BaseInfoList = v.BaseInfoList[:0]
	poolBaseInfosDto.Put(v)
}
