package scbp

// AdsTargetingTagDto 
type AdsTargetingTagDto struct {
    // 定向标签id
    TagId   int64 `json:"tag_id,omitempty" xml:"tag_id,omitempty"`
    // 计划id
    CampaignId   int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
    // 溢价
    Discount   int64 `json:"discount,omitempty" xml:"discount,omitempty"`
    // 标签选项
    OptionValue   string `json:"option_value,omitempty" xml:"option_value,omitempty"`
    // 人群名称
    CrowdName   string `json:"crowd_name,omitempty" xml:"crowd_name,omitempty"`
    // 效果数据
    Effect   *AdsEffectDto `json:"effect,omitempty" xml:"effect,omitempty"`
    // 出价类型：0=出价, 1=溢价，2=过滤, 3=召回
    PriceMode   int64 `json:"price_mode,omitempty" xml:"price_mode,omitempty"`
    // 产品线id
    ProductLineId   int64 `json:"product_line_id,omitempty" xml:"product_line_id,omitempty"`
    // 计划创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 计划修改时间
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 13:地域标签 14：人群标签
    TagRefType   int64 `json:"tag_ref_type,omitempty" xml:"tag_ref_type,omitempty"`
}
