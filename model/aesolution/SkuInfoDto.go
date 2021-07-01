package aesolution

// SkuInfoDto 结构体
type SkuInfoDto struct {
	// sku attribute list. Some categories don't have sku attributes, then sku_attributes_list should be empty.
	SkuAttributesList []SkuAttributeDto `json:"sku_attributes_list,omitempty" xml:"sku_attributes_list>sku_attribute_dto,omitempty"`
	// extra params. Configured some special products
	ExtraParams string `json:"extra_params,omitempty" xml:"extra_params,omitempty"`
	// stock. Maximum:999999, minumum:1
	Inventory int64 `json:"inventory,omitempty" xml:"inventory,omitempty"`
	// price. Maximum:999999, minumum:0.01
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// Code for merchant's sku, important reference to maintain the relationship between merchant and Aliexpress's system.
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// discount price for the sku. discount_price should be cheaper than price.
	DiscountPrice string `json:"discount_price,omitempty" xml:"discount_price,omitempty"`
	// barcode of the sku. Except some Russian sellers who will be fulfilled by Aliexpress, please ignore this field for other sellers.
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
}
