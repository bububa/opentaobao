package alsc

// QueryCardTemplateOpenReq 
type QueryCardTemplateOpenReq struct {
    // 品牌id
    BrandId   string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    // 卡模板ID
    CardTemplateId   string `json:"card_template_id,omitempty" xml:"card_template_id,omitempty"`
    // 外部品牌id,brandId与out_brand_id不可同时为空
    OutBrandId   string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
    // 外部门店id
    OutShopId   string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
    // 门店id
    ShopId   string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
    // 是否包含逻辑删除
    IncludeLogicalDelete   bool `json:"include_logical_delete,omitempty" xml:"include_logical_delete,omitempty"`
    // 最后一次更新时间
    MaxUpdateTime   string `json:"max_update_time,omitempty" xml:"max_update_time,omitempty"`
    // 物理卡号
    PhysicalCardId   string `json:"physical_card_id,omitempty" xml:"physical_card_id,omitempty"`
}
