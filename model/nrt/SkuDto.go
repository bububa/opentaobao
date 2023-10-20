package nrt

import (
	"sync"
)

// SkuDto 结构体
type SkuDto struct {
	// SKU属性
	Properties []CategoryPropDto `json:"properties,omitempty" xml:"properties>category_prop_dto,omitempty"`
	// 条码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 商家编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 修改时间
	UpdateTime string `json:"update_time,omitempty" xml:"update_time,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// SKU ID
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolSkuDto = sync.Pool{
	New: func() any {
		return new(SkuDto)
	},
}

// GetSkuDto() 从对象池中获取SkuDto
func GetSkuDto() *SkuDto {
	return poolSkuDto.Get().(*SkuDto)
}

// ReleaseSkuDto 释放SkuDto
func ReleaseSkuDto(v *SkuDto) {
	v.Properties = v.Properties[:0]
	v.Barcode = ""
	v.CreateTime = ""
	v.OuterId = ""
	v.Price = ""
	v.UpdateTime = ""
	v.Quantity = 0
	v.SkuId = 0
	poolSkuDto.Put(v)
}
