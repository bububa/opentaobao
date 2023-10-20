package feedflow

import (
	"sync"
)

// ErrorInfoDto 结构体
type ErrorInfoDto struct {
	// 该原因失败对象列表
	ErrorObjectList []ErrorObjectDto `json:"error_object_list,omitempty" xml:"error_object_list>error_object_dto,omitempty"`
	// 失败原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
}

var poolErrorInfoDto = sync.Pool{
	New: func() any {
		return new(ErrorInfoDto)
	},
}

// GetErrorInfoDto() 从对象池中获取ErrorInfoDto
func GetErrorInfoDto() *ErrorInfoDto {
	return poolErrorInfoDto.Get().(*ErrorInfoDto)
}

// ReleaseErrorInfoDto 释放ErrorInfoDto
func ReleaseErrorInfoDto(v *ErrorInfoDto) {
	v.ErrorObjectList = v.ErrorObjectList[:0]
	v.Reason = ""
	poolErrorInfoDto.Put(v)
}
