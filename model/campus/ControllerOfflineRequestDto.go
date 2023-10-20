package campus

import (
	"sync"
)

// ControllerOfflineRequestDto 结构体
type ControllerOfflineRequestDto struct {
	// sn
	SnNo string `json:"sn_no,omitempty" xml:"sn_no,omitempty"`
	// 离线日志JSON
	OfflineLog string `json:"offline_log,omitempty" xml:"offline_log,omitempty"`
}

var poolControllerOfflineRequestDto = sync.Pool{
	New: func() any {
		return new(ControllerOfflineRequestDto)
	},
}

// GetControllerOfflineRequestDto() 从对象池中获取ControllerOfflineRequestDto
func GetControllerOfflineRequestDto() *ControllerOfflineRequestDto {
	return poolControllerOfflineRequestDto.Get().(*ControllerOfflineRequestDto)
}

// ReleaseControllerOfflineRequestDto 释放ControllerOfflineRequestDto
func ReleaseControllerOfflineRequestDto(v *ControllerOfflineRequestDto) {
	v.SnNo = ""
	v.OfflineLog = ""
	poolControllerOfflineRequestDto.Put(v)
}
