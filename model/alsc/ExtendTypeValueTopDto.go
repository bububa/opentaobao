package alsc

import (
	"sync"
)

// ExtendTypeValueTopDto 结构体
type ExtendTypeValueTopDto struct {
	// 扩展code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 扩展value
	Obj string `json:"obj,omitempty" xml:"obj,omitempty"`
}

var poolExtendTypeValueTopDto = sync.Pool{
	New: func() any {
		return new(ExtendTypeValueTopDto)
	},
}

// GetExtendTypeValueTopDto() 从对象池中获取ExtendTypeValueTopDto
func GetExtendTypeValueTopDto() *ExtendTypeValueTopDto {
	return poolExtendTypeValueTopDto.Get().(*ExtendTypeValueTopDto)
}

// ReleaseExtendTypeValueTopDto 释放ExtendTypeValueTopDto
func ReleaseExtendTypeValueTopDto(v *ExtendTypeValueTopDto) {
	v.Code = ""
	v.Obj = ""
	poolExtendTypeValueTopDto.Put(v)
}
