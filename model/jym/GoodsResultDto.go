package jym

import (
	"sync"
)

// GoodsResultDto 结构体
type GoodsResultDto struct {
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 状态码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否下架成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolGoodsResultDto = sync.Pool{
	New: func() any {
		return new(GoodsResultDto)
	},
}

// GetGoodsResultDto() 从对象池中获取GoodsResultDto
func GetGoodsResultDto() *GoodsResultDto {
	return poolGoodsResultDto.Get().(*GoodsResultDto)
}

// ReleaseGoodsResultDto 释放GoodsResultDto
func ReleaseGoodsResultDto(v *GoodsResultDto) {
	v.Desc = ""
	v.Code = 0
	v.Success = false
	poolGoodsResultDto.Put(v)
}
