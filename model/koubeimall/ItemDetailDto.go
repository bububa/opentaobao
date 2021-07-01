package koubeimall

// ItemDetailDto 
type ItemDetailDto struct {
    // 商品使用规则
    ItemRule   *ItemRule `json:"item_rule,omitempty" xml:"item_rule,omitempty"`
    // 购买须知
    ItemBuyNotes   *ItemBuyNotes `json:"item_buy_notes,omitempty" xml:"item_buy_notes,omitempty"`
    // 详情内容
    ItemGroupDetailList   []ItemGroupContent `json:"item_group_detail_list,omitempty" xml:"item_group_detail_list>item_group_content,omitempty"`
    // 图文详情
    ItemImageDetailList   []ItemImageDetail `json:"item_image_detail_list,omitempty" xml:"item_image_detail_list>item_image_detail,omitempty"`
    // 商品服务
    ItemServices   *ItemServices `json:"item_services,omitempty" xml:"item_services,omitempty"`
    // 补充说明
    MemoList   []string `json:"memo_list,omitempty" xml:"memo_list>string,omitempty"`
}
