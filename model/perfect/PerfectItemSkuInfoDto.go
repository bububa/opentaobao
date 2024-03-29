package perfect

import (
	"sync"
)

// PerfectItemSkuInfoDto 结构体
type PerfectItemSkuInfoDto struct {
	// sku销售属性
	SaleProperties []PerfectSalePropertyDto `json:"sale_properties,omitempty" xml:"sale_properties>perfect_sale_property_dto,omitempty"`
	// sku条码
	SkuBarcode string `json:"sku_barcode,omitempty" xml:"sku_barcode,omitempty"`
	// sku外部编码
	SkuOuterId string `json:"sku_outer_id,omitempty" xml:"sku_outer_id,omitempty"`
	// sku吊牌价
	SkuPretium string `json:"sku_pretium,omitempty" xml:"sku_pretium,omitempty"`
	// sku价格
	SkuPrice string `json:"sku_price,omitempty" xml:"sku_price,omitempty"`
	// sku关联货品
	ScProductInfo *PerfectScProductInfoDto `json:"sc_product_info,omitempty" xml:"sc_product_info,omitempty"`
	// sku数量
	SkuQuantity int64 `json:"sku_quantity,omitempty" xml:"sku_quantity,omitempty"`
}

var poolPerfectItemSkuInfoDto = sync.Pool{
	New: func() any {
		return new(PerfectItemSkuInfoDto)
	},
}

// GetPerfectItemSkuInfoDto() 从对象池中获取PerfectItemSkuInfoDto
func GetPerfectItemSkuInfoDto() *PerfectItemSkuInfoDto {
	return poolPerfectItemSkuInfoDto.Get().(*PerfectItemSkuInfoDto)
}

// ReleasePerfectItemSkuInfoDto 释放PerfectItemSkuInfoDto
func ReleasePerfectItemSkuInfoDto(v *PerfectItemSkuInfoDto) {
	v.SaleProperties = v.SaleProperties[:0]
	v.SkuBarcode = ""
	v.SkuOuterId = ""
	v.SkuPretium = ""
	v.SkuPrice = ""
	v.ScProductInfo = nil
	v.SkuQuantity = 0
	poolPerfectItemSkuInfoDto.Put(v)
}
