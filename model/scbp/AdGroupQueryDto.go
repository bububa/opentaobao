package scbp

// AdGroupQueryDto 
type AdGroupQueryDto struct {
    // adgroup上下线状态
    AdgroupOnlineStatus   int64 `json:"adgroup_online_status,omitempty" xml:"adgroup_online_status,omitempty"`
    // FEED上下线状态
    FeedOnlineStatus   int64 `json:"feed_online_status,omitempty" xml:"feed_online_status,omitempty"`
    // 最小曝光数量
    MinImprCnt   int64 `json:"min_impr_cnt,omitempty" xml:"min_impr_cnt,omitempty"`
    // 新品成长 如果是1，则是爆品潜力品
    HotPotentialProduct   int64 `json:"hot_potential_product,omitempty" xml:"hot_potential_product,omitempty"`
    // 产品id
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
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
    // 计划类型
    CampaignType   int64 `json:"campaign_type,omitempty" xml:"campaign_type,omitempty"`
    // 计划状态
    CampaignStatus   int64 `json:"campaign_status,omitempty" xml:"campaign_status,omitempty"`
    // 推广组id
    AdGroupId   int64 `json:"ad_group_id,omitempty" xml:"ad_group_id,omitempty"`
    // 推广组状态
    AdGroupStatus   int64 `json:"ad_group_status,omitempty" xml:"ad_group_status,omitempty"`
    // 产品线id
    AdsLineId   int64 `json:"ads_line_id,omitempty" xml:"ads_line_id,omitempty"`
    // 产品id
    AdsId   int64 `json:"ads_id,omitempty" xml:"ads_id,omitempty"`
    // 标题
    Subject   string `json:"subject,omitempty" xml:"subject,omitempty"`
    // 发品用户id
    OwnerMemberId   string `json:"owner_member_id,omitempty" xml:"owner_member_id,omitempty"`
    // 排序字段
    OrderBy   string `json:"order_by,omitempty" xml:"order_by,omitempty"`
    // 正序
    Order   string `json:"order,omitempty" xml:"order,omitempty"`
    // 页码
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`
    // 每页数量
    Size   int64 `json:"size,omitempty" xml:"size,omitempty"`
    // 分页
    Paging   bool `json:"paging,omitempty" xml:"paging,omitempty"`
}
