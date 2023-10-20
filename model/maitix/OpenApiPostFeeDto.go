package maitix

import (
	"sync"
)

// OpenApiPostFeeDto 结构体
type OpenApiPostFeeDto struct {
	// 运费金额，单位分
	PostFee int64 `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
}

var poolOpenApiPostFeeDto = sync.Pool{
	New: func() any {
		return new(OpenApiPostFeeDto)
	},
}

// GetOpenApiPostFeeDto() 从对象池中获取OpenApiPostFeeDto
func GetOpenApiPostFeeDto() *OpenApiPostFeeDto {
	return poolOpenApiPostFeeDto.Get().(*OpenApiPostFeeDto)
}

// ReleaseOpenApiPostFeeDto 释放OpenApiPostFeeDto
func ReleaseOpenApiPostFeeDto(v *OpenApiPostFeeDto) {
	v.PostFee = 0
	poolOpenApiPostFeeDto.Put(v)
}
