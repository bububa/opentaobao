package product

import (
	"sync"
)

// GoodsBaseInfoDto 结构体
type GoodsBaseInfoDto struct {
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 商品描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 库存
	Storage int64 `json:"storage,omitempty" xml:"storage,omitempty"`
}

var poolGoodsBaseInfoDto = sync.Pool{
	New: func() any {
		return new(GoodsBaseInfoDto)
	},
}

// GetGoodsBaseInfoDto() 从对象池中获取GoodsBaseInfoDto
func GetGoodsBaseInfoDto() *GoodsBaseInfoDto {
	return poolGoodsBaseInfoDto.Get().(*GoodsBaseInfoDto)
}

// ReleaseGoodsBaseInfoDto 释放GoodsBaseInfoDto
func ReleaseGoodsBaseInfoDto(v *GoodsBaseInfoDto) {
	v.Title = ""
	v.Price = ""
	v.Description = ""
	v.Storage = 0
	poolGoodsBaseInfoDto.Put(v)
}
