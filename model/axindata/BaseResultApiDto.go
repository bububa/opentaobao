package axindata

import (
	"sync"
)

// BaseResultApiDto 结构体
type BaseResultApiDto struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误code
	ErCode string `json:"er_code,omitempty" xml:"er_code,omitempty"`
	// 错误msg
	ErMsg string `json:"er_msg,omitempty" xml:"er_msg,omitempty"`
	// 额外信息
	InfoMsg string `json:"info_msg,omitempty" xml:"info_msg,omitempty"`
	// 数据模型
	Data *StdHotelDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 调用是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}

var poolBaseResultApiDto = sync.Pool{
	New: func() any {
		return new(BaseResultApiDto)
	},
}

// GetBaseResultApiDto() 从对象池中获取BaseResultApiDto
func GetBaseResultApiDto() *BaseResultApiDto {
	return poolBaseResultApiDto.Get().(*BaseResultApiDto)
}

// ReleaseBaseResultApiDto 释放BaseResultApiDto
func ReleaseBaseResultApiDto(v *BaseResultApiDto) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.ErCode = ""
	v.ErMsg = ""
	v.InfoMsg = ""
	v.Data = nil
	v.Success = false
	v.Succ = false
	poolBaseResultApiDto.Put(v)
}
