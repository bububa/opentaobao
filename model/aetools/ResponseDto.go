package aetools

import (
	"sync"
)

// ResponseDto 结构体
type ResponseDto struct {
	// 返回结果状态描述
	RespMsg string `json:"resp_msg,omitempty" xml:"resp_msg,omitempty"`
	// 返回结果状态码
	RespCode int64 `json:"resp_code,omitempty" xml:"resp_code,omitempty"`
	// 返回结果明细信息
	Result *PromotionLinkResultDto `json:"result,omitempty" xml:"result,omitempty"`
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
	v.RespMsg = ""
	v.RespCode = 0
	v.Result = nil
	poolResponseDto.Put(v)
}
