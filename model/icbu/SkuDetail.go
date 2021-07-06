package icbu

// SkuDetail 结构体
type SkuDetail struct {
	// 商品属性
	Attributes []ProductAttribute `json:"attributes,omitempty" xml:"attributes>product_attribute,omitempty"`
	// 库存
	InventoryDtoList []InventoryDetail `json:"inventory_dto_list,omitempty" xml:"inventory_dto_list>inventory_detail,omitempty"`
	// 价格，单位是美元，精确到小数点后两位，范围是0.01-9999999.00
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 商品的SKU编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// SKU id，唯一标识一个SKU
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
