package aliexpresssumaitong

import (
	"sync"
)

// ResponseDto 结构体
type ResponseDto struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 数据
	Data *HjTaxCalculateResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 成功标记
	Succeeded bool `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
}

var poolResponseDto = sync.Pool{
	New: func() any {
		return new(ResponseDto)
	},
}

// GetResponseDto() 从对象池中获取ResponseDto
func GetResponseDto() *ResponseDto {
	return poolResponseDto.Get().(*ResponseDto)
}

// ReleaseResponseDto 释放ResponseDto
func ReleaseResponseDto(v *ResponseDto) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Succeeded = false
	poolResponseDto.Put(v)
}
