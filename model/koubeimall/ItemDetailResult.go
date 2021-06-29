package koubeimall

// ItemDetailResult 
type ItemDetailResult struct {
    // 商品详情信息
    ItemDetail   *ItemDetailDTO `json:"item_detail,omitempty" xml:"item_detail,omitempty"`
    // 商品基础信息
    ItemInfo   *ItemDTO `json:"item_info,omitempty" xml:"item_info,omitempty"`
}
