package dmp

import (
	"sync"
)

// ApiContextDto 结构体
type ApiContextDto struct {
	// 业务线编码，引力魔方：displayDefault
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
}

var poolApiContextDto = sync.Pool{
	New: func() any {
		return new(ApiContextDto)
	},
}

// GetApiContextDto() 从对象池中获取ApiContextDto
func GetApiContextDto() *ApiContextDto {
	return poolApiContextDto.Get().(*ApiContextDto)
}

// ReleaseApiContextDto 释放ApiContextDto
func ReleaseApiContextDto(v *ApiContextDto) {
	v.BizCode = ""
	poolApiContextDto.Put(v)
}
