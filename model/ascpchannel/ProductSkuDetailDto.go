package ascpchannel

// ProductSkuDetailDto 结构体
type ProductSkuDetailDto struct {
	// sku 销售属性
	Properties []string `json:"properties,omitempty" xml:"properties>string,omitempty"`
	// sku图片链接
	Picture []string `json:"picture,omitempty" xml:"picture>string,omitempty"`
	// skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
