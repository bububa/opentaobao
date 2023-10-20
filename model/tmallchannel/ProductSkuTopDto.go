package tmallchannel

import (
	"sync"
)

// ProductSkuTopDto 结构体
type ProductSkuTopDto struct {
	// sku商家编码
	SkuNumber string `json:"sku_number,omitempty" xml:"sku_number,omitempty"`
	// 条形码
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 图片链接
	PictureUrl string `json:"picture_url,omitempty" xml:"picture_url,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// sku后端货品
	SkuScItemId int64 `json:"sku_sc_item_id,omitempty" xml:"sku_sc_item_id,omitempty"`
	// 基准价
	StandardPrice int64 `json:"standard_price,omitempty" xml:"standard_price,omitempty"`
}

var poolProductSkuTopDto = sync.Pool{
	New: func() any {
		return new(ProductSkuTopDto)
	},
}

// GetProductSkuTopDto() 从对象池中获取ProductSkuTopDto
func GetProductSkuTopDto() *ProductSkuTopDto {
	return poolProductSkuTopDto.Get().(*ProductSkuTopDto)
}

// ReleaseProductSkuTopDto 释放ProductSkuTopDto
func ReleaseProductSkuTopDto(v *ProductSkuTopDto) {
	v.SkuNumber = ""
	v.BarCode = ""
	v.PictureUrl = ""
	v.SkuId = 0
	v.SkuScItemId = 0
	v.StandardPrice = 0
	poolProductSkuTopDto.Put(v)
}
