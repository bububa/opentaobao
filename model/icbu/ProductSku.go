package icbu

// ProductSku 结构体
type ProductSku struct {
	// 商品属性
	Attributes []ProductAttribute `json:"attributes,omitempty" xml:"attributes>product_attribute,omitempty"`
	// 需要失效的SKU的详细定义
	ExcludeSkus []SkuDetail `json:"exclude_skus,omitempty" xml:"exclude_skus>sku_detail,omitempty"`
	// 单个SKU详细定义
	SpecialSkus []SkuDetail `json:"special_skus,omitempty" xml:"special_skus>sku_detail,omitempty"`
}
