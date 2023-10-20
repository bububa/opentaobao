package tblogistics

import (
	"sync"
)

// WriteOffGoodsDto 结构体
type WriteOffGoodsDto struct {
	// 商品Id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 名称
	GoodsName string `json:"goods_name,omitempty" xml:"goods_name,omitempty"`
	// 商品sku id
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 商品图片地址
	GoodsPicId string `json:"goods_pic_id,omitempty" xml:"goods_pic_id,omitempty"`
	// 数量
	GoodsQuantity int64 `json:"goods_quantity,omitempty" xml:"goods_quantity,omitempty"`
}

var poolWriteOffGoodsDto = sync.Pool{
	New: func() any {
		return new(WriteOffGoodsDto)
	},
}

// GetWriteOffGoodsDto() 从对象池中获取WriteOffGoodsDto
func GetWriteOffGoodsDto() *WriteOffGoodsDto {
	return poolWriteOffGoodsDto.Get().(*WriteOffGoodsDto)
}

// ReleaseWriteOffGoodsDto 释放WriteOffGoodsDto
func ReleaseWriteOffGoodsDto(v *WriteOffGoodsDto) {
	v.ItemId = ""
	v.Price = ""
	v.GoodsName = ""
	v.SkuId = ""
	v.GoodsPicId = ""
	v.GoodsQuantity = 0
	poolWriteOffGoodsDto.Put(v)
}
