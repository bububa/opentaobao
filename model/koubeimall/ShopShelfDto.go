package koubeimall

// ShopShelfDto 
type ShopShelfDto struct {
    // 是否有更多商品
    HasMore   bool `json:"has_more,omitempty" xml:"has_more,omitempty"`
    // 每页查询长度
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 货架商品总数
    TotalItemSize   int64 `json:"total_item_size,omitempty" xml:"total_item_size,omitempty"`
    // 下一页开始起始值
    NextStart   int64 `json:"next_start,omitempty" xml:"next_start,omitempty"`
    // 商品模型
    ItemInfoList   []ItemDto `json:"item_info_list,omitempty" xml:"item_info_list>item_dto,omitempty"`
}
