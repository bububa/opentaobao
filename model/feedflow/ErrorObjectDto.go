package feedflow

import (
	"sync"
)

// ErrorObjectDto 结构体
type ErrorObjectDto struct {
	// 创意id
	CreativeId int64 `json:"creative_id,omitempty" xml:"creative_id,omitempty"`
}

var poolErrorObjectDto = sync.Pool{
	New: func() any {
		return new(ErrorObjectDto)
	},
}

// GetErrorObjectDto() 从对象池中获取ErrorObjectDto
func GetErrorObjectDto() *ErrorObjectDto {
	return poolErrorObjectDto.Get().(*ErrorObjectDto)
}

// ReleaseErrorObjectDto 释放ErrorObjectDto
func ReleaseErrorObjectDto(v *ErrorObjectDto) {
	v.CreativeId = 0
	poolErrorObjectDto.Put(v)
}
