package tmallcar

// Item4isvDto 结构体
type Item4isvDto struct {
	// 附加属性
	Extension string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 商品缩略图
	ItemHeadImg string `json:"item_head_img,omitempty" xml:"item_head_img,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 商品服务价格,单位：元
	ServicePrice string `json:"service_price,omitempty" xml:"service_price,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
