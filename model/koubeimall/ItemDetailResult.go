package koubeimall

// ItemDetailResult 结构体
type ItemDetailResult struct {
	// 商品详情信息
	ItemDetail *ItemDetailDto `json:"item_detail,omitempty" xml:"item_detail,omitempty"`
	// 商品基础信息
	ItemInfo *ItemDto `json:"item_info,omitempty" xml:"item_info,omitempty"`
}
