package campus

import (
	"sync"
)

// ControllerOfflineDataDto 结构体
type ControllerOfflineDataDto struct {
	// 门禁点数据
	GuardOfflineList []GuardOfflineDataDto `json:"guard_offline_list,omitempty" xml:"guard_offline_list>guard_offline_data_dto,omitempty"`
}

var poolControllerOfflineDataDto = sync.Pool{
	New: func() any {
		return new(ControllerOfflineDataDto)
	},
}

// GetControllerOfflineDataDto() 从对象池中获取ControllerOfflineDataDto
func GetControllerOfflineDataDto() *ControllerOfflineDataDto {
	return poolControllerOfflineDataDto.Get().(*ControllerOfflineDataDto)
}

// ReleaseControllerOfflineDataDto 释放ControllerOfflineDataDto
func ReleaseControllerOfflineDataDto(v *ControllerOfflineDataDto) {
	v.GuardOfflineList = v.GuardOfflineList[:0]
	poolControllerOfflineDataDto.Put(v)
}
