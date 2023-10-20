package tmallchannel

import (
	"sync"
)

// ProductTopDto 结构体
type ProductTopDto struct {
	// sku列表
	SkuList []ProductSkuTopDto `json:"sku_list,omitempty" xml:"sku_list>product_sku_top_dto,omitempty"`
	// 产品编码
	ProductNumber string `json:"product_number,omitempty" xml:"product_number,omitempty"`
	// 产品描述地址
	DescPath string `json:"desc_path,omitempty" xml:"desc_path,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 产品线ID
	ProductLineId int64 `json:"product_line_id,omitempty" xml:"product_line_id,omitempty"`
	// 基准价
	StandardPrice int64 `json:"standard_price,omitempty" xml:"standard_price,omitempty"`
	// 产品Id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 没有sku的情况下，产品对应的后端商品id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// spuId
	SpuId int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	// 供应商Id
	SupplierId int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 类目Id
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
}

var poolProductTopDto = sync.Pool{
	New: func() any {
		return new(ProductTopDto)
	},
}

// GetProductTopDto() 从对象池中获取ProductTopDto
func GetProductTopDto() *ProductTopDto {
	return poolProductTopDto.Get().(*ProductTopDto)
}

// ReleaseProductTopDto 释放ProductTopDto
func ReleaseProductTopDto(v *ProductTopDto) {
	v.SkuList = v.SkuList[:0]
	v.ProductNumber = ""
	v.DescPath = ""
	v.Title = ""
	v.ProductLineId = 0
	v.StandardPrice = 0
	v.ProductId = 0
	v.ScItemId = 0
	v.SpuId = 0
	v.SupplierId = 0
	v.CategoryId = 0
	poolProductTopDto.Put(v)
}
