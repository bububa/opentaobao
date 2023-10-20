package axintrade

import (
	"sync"
)

// BaseResultApiDto 结构体
type BaseResultApiDto struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回素材id
	Data *HotelOrderRefundResApiDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
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
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Success = false
	poolBaseResultApiDto.Put(v)
}
