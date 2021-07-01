package product

// UpdateSkuPrice 结构体
type UpdateSkuPrice struct {
	// SkuID，用于指定被修改价格的SKU
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充，用于指定被修改价格的SKU
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
	// Sku的商家外部id，用于指定被修改价格的SKU
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 属于这个sku的商品的价格 取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。
	Price string `json:"price,omitempty" xml:"price,omitempty"`
}
