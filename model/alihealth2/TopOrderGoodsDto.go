package alihealth2

import (
	"sync"
)

// TopOrderGoodsDto 结构体
type TopOrderGoodsDto struct {
	// 商品编码
	GoodsCode string `json:"goods_code,omitempty" xml:"goods_code,omitempty"`
	// 商品名称
	GoodsName string `json:"goods_name,omitempty" xml:"goods_name,omitempty"`
	// 商品数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}

var poolTopOrderGoodsDto = sync.Pool{
	New: func() any {
		return new(TopOrderGoodsDto)
	},
}

// GetTopOrderGoodsDto() 从对象池中获取TopOrderGoodsDto
func GetTopOrderGoodsDto() *TopOrderGoodsDto {
	return poolTopOrderGoodsDto.Get().(*TopOrderGoodsDto)
}

// ReleaseTopOrderGoodsDto 释放TopOrderGoodsDto
func ReleaseTopOrderGoodsDto(v *TopOrderGoodsDto) {
	v.GoodsCode = ""
	v.GoodsName = ""
	v.Count = 0
	poolTopOrderGoodsDto.Put(v)
}
