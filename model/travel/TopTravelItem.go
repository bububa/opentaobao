package travel

// TopTravelItem 结构体
type TopTravelItem struct {
	// 商家自定义商品编码
	OutProductId string `json:"out_product_id,omitempty" xml:"out_product_id,omitempty"`
	// 商品修改时间
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 扩展信息
	Extend string `json:"extend,omitempty" xml:"extend,omitempty"`
	// 商品创建时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 商家编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// skuid
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
