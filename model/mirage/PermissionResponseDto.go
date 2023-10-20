package mirage

import (
	"sync"
)

// PermissionResponseDto 结构体
type PermissionResponseDto struct {
	// 错误内容
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 该值为MAP类型，KEY=视频或者节目ID；VALUE中包含两个值resourceId视频或者节目ID(String类型)，以及access该视频或者节目是否可播（Boolean）结果
	Permissions string `json:"permissions,omitempty" xml:"permissions,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolPermissionResponseDto = sync.Pool{
	New: func() any {
		return new(PermissionResponseDto)
	},
}

// GetPermissionResponseDto() 从对象池中获取PermissionResponseDto
func GetPermissionResponseDto() *PermissionResponseDto {
	return poolPermissionResponseDto.Get().(*PermissionResponseDto)
}

// ReleasePermissionResponseDto 释放PermissionResponseDto
func ReleasePermissionResponseDto(v *PermissionResponseDto) {
	v.Message = ""
	v.Permissions = ""
	v.Success = false
	poolPermissionResponseDto.Put(v)
}
