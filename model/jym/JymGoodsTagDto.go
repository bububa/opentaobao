package jym

import (
	"sync"
)

// JymGoodsTagDto 结构体
type JymGoodsTagDto struct {
	// 标签名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 标签类型，1-服务 2-卖点
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolJymGoodsTagDto = sync.Pool{
	New: func() any {
		return new(JymGoodsTagDto)
	},
}

// GetJymGoodsTagDto() 从对象池中获取JymGoodsTagDto
func GetJymGoodsTagDto() *JymGoodsTagDto {
	return poolJymGoodsTagDto.Get().(*JymGoodsTagDto)
}

// ReleaseJymGoodsTagDto 释放JymGoodsTagDto
func ReleaseJymGoodsTagDto(v *JymGoodsTagDto) {
	v.Name = ""
	v.Type = 0
	poolJymGoodsTagDto.Put(v)
}
