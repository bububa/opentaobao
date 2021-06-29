package alsc

// PlanRuleQueryOpenReq 
type PlanRuleQueryOpenReq struct {
    // 是否包含逻辑删除的数据
    Deleted   bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
    // 最后更新时间
    LastUpdateTime   string `json:"last_update_time,omitempty" xml:"last_update_time,omitempty"`
    // 品牌ID
    BrandId   string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    // 外部品牌id
    OutBrandId   string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
    // 每页数量
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 上次最有一个ID
    LastId   string `json:"last_id,omitempty" xml:"last_id,omitempty"`
    // 店铺ID
    ShopId   string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
    // 外部店铺ID
    OutShopId   string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
}
