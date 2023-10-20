package omniorder

import (
	"sync"
)

// ItemLightPublishSkuDto 结构体
type ItemLightPublishSkuDto struct {
	// salePropsDTO
	SalePropsDTOs []ItemLightPublishSalePropDto `json:"sale_props_d_t_os,omitempty" xml:"sale_props_d_t_os>item_light_publish_sale_prop_dto,omitempty"`
	// sku图片
	SkuImages []string `json:"sku_images,omitempty" xml:"sku_images>string,omitempty"`
	// sku销售属性
	SaleProps []ItemLightPublishSalePropDto `json:"sale_props,omitempty" xml:"sale_props>item_light_publish_sale_prop_dto,omitempty"`
	// sku条形码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// sku扩展字段
	ExtendAttr string `json:"extend_attr,omitempty" xml:"extend_attr,omitempty"`
	// sku销售价
	Pretium string `json:"pretium,omitempty" xml:"pretium,omitempty"`
	// sku吊牌价
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// skuOuterId
	SkuOuterId string `json:"sku_outer_id,omitempty" xml:"sku_outer_id,omitempty"`
	// sku条形码
	SkuBarcode string `json:"sku_barcode,omitempty" xml:"sku_barcode,omitempty"`
	// customCode
	CustomCode string `json:"custom_code,omitempty" xml:"custom_code,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolItemLightPublishSkuDto = sync.Pool{
	New: func() any {
		return new(ItemLightPublishSkuDto)
	},
}

// GetItemLightPublishSkuDto() 从对象池中获取ItemLightPublishSkuDto
func GetItemLightPublishSkuDto() *ItemLightPublishSkuDto {
	return poolItemLightPublishSkuDto.Get().(*ItemLightPublishSkuDto)
}

// ReleaseItemLightPublishSkuDto 释放ItemLightPublishSkuDto
func ReleaseItemLightPublishSkuDto(v *ItemLightPublishSkuDto) {
	v.SalePropsDTOs = v.SalePropsDTOs[:0]
	v.SkuImages = v.SkuImages[:0]
	v.SaleProps = v.SaleProps[:0]
	v.Barcode = ""
	v.ExtendAttr = ""
	v.Pretium = ""
	v.Price = ""
	v.SkuOuterId = ""
	v.SkuBarcode = ""
	v.CustomCode = ""
	v.SkuId = 0
	poolItemLightPublishSkuDto.Put(v)
}
