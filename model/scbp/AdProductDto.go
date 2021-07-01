package scbp

// AdProductDto 
type AdProductDto struct {
    // 计划id
    CampaignId   int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
    // 计划类型
    CampaignType   int64 `json:"campaign_type,omitempty" xml:"campaign_type,omitempty"`
    // 计划状态
    CampaignStatus   int64 `json:"campaign_status,omitempty" xml:"campaign_status,omitempty"`
    // 组id
    AdGroupId   int64 `json:"ad_group_id,omitempty" xml:"ad_group_id,omitempty"`
    // 推广组状态
    AdGroupStatus   int64 `json:"ad_group_status,omitempty" xml:"ad_group_status,omitempty"`
    // 产品线id
    AdsLineId   int64 `json:"ads_line_id,omitempty" xml:"ads_line_id,omitempty"`
    // 产品id
    AdsId   int64 `json:"ads_id,omitempty" xml:"ads_id,omitempty"`
    // 主键id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 商品id
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
    // 商品标题
    Subject   string `json:"subject,omitempty" xml:"subject,omitempty"`
    // 1级分组
    GroupLv1Id   int64 `json:"group_lv1_id,omitempty" xml:"group_lv1_id,omitempty"`
    // 2级分组
    GroupLv2Id   int64 `json:"group_lv2_id,omitempty" xml:"group_lv2_id,omitempty"`
    // 3级分组
    GroupLv3Id   int64 `json:"group_lv3_id,omitempty" xml:"group_lv3_id,omitempty"`
    // 1级类目
    CateLv1Id   int64 `json:"cate_lv1_id,omitempty" xml:"cate_lv1_id,omitempty"`
    // 2级类目
    CateLv2Id   int64 `json:"cate_lv2_id,omitempty" xml:"cate_lv2_id,omitempty"`
    // 3级类目
    CateLv3Id   int64 `json:"cate_lv3_id,omitempty" xml:"cate_lv3_id,omitempty"`
    // posting创建时间
    GmtPostingCreate   string `json:"gmt_posting_create,omitempty" xml:"gmt_posting_create,omitempty"`
    // posting修改时间
    GmtPostingModified   string `json:"gmt_posting_modified,omitempty" xml:"gmt_posting_modified,omitempty"`
    // 效果数据
    Effect   *AdProductEffectDto `json:"effect,omitempty" xml:"effect,omitempty"`
}
