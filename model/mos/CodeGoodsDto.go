package mos

import (
	"sync"
)

// CodeGoodsDto 结构体
type CodeGoodsDto struct {
	// zi单号
	SubNo string `json:"sub_no,omitempty" xml:"sub_no,omitempty"`
	// 数量
	Quantity float64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// id
	GoodsId int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
}

var poolCodeGoodsDto = sync.Pool{
	New: func() any {
		return new(CodeGoodsDto)
	},
}

// GetCodeGoodsDto() 从对象池中获取CodeGoodsDto
func GetCodeGoodsDto() *CodeGoodsDto {
	return poolCodeGoodsDto.Get().(*CodeGoodsDto)
}

// ReleaseCodeGoodsDto 释放CodeGoodsDto
func ReleaseCodeGoodsDto(v *CodeGoodsDto) {
	v.SubNo = ""
	v.Quantity = 0
	v.GoodsId = 0
	poolCodeGoodsDto.Put(v)
}
