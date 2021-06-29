package feedflow

// RptResultDTO 
type RptResultDTO struct {
    // 消耗
    Charge   string `json:"charge,omitempty" xml:"charge,omitempty"`
    // 有效展现
    AdPv   int64 `json:"ad_pv,omitempty" xml:"ad_pv,omitempty"`
    // 有效点击数
    Click   int64 `json:"click,omitempty" xml:"click,omitempty"`
    // 千次展现成本
    Ecpm   string `json:"ecpm,omitempty" xml:"ecpm,omitempty"`
    // 点击单价
    Ecpc   string `json:"ecpc,omitempty" xml:"ecpc,omitempty"`
    // 引导访问量
    InshopPv   int64 `json:"inshop_pv,omitempty" xml:"inshop_pv,omitempty"`
    // 引导进店人数
    InshopUv   int64 `json:"inshop_uv,omitempty" xml:"inshop_uv,omitempty"`
    // 引导进店率
    InshopUvRate   string `json:"inshop_uv_rate,omitempty" xml:"inshop_uv_rate,omitempty"`
    // 深度进店次数
    DeepInshopNum   int64 `json:"deep_inshop_num,omitempty" xml:"deep_inshop_num,omitempty"`
    // 平均访问时长
    AvgAccessTime   string `json:"avg_access_time,omitempty" xml:"avg_access_time,omitempty"`
    // 平均访问页面数
    AvgAccessPageNum   string `json:"avg_access_page_num,omitempty" xml:"avg_access_page_num,omitempty"`
    // 粉丝关注量
    FollowNumber   int64 `json:"follow_number,omitempty" xml:"follow_number,omitempty"`
    // 新客获取量
    AddNewUv   int64 `json:"add_new_uv,omitempty" xml:"add_new_uv,omitempty"`
    // 新客获取率
    AddNewUvRate   string `json:"add_new_uv_rate,omitempty" xml:"add_new_uv_rate,omitempty"`
    // 拉新消耗
    NewFCharge   string `json:"new_f_charge,omitempty" xml:"new_f_charge,omitempty"`
    // 收藏宝贝数
    InshopItemColNum   int64 `json:"inshop_item_col_num,omitempty" xml:"inshop_item_col_num,omitempty"`
    // 加购物车数
    CartNum   int64 `json:"cart_num,omitempty" xml:"cart_num,omitempty"`
    // 拍下订单量
    GmvInshopNum   int64 `json:"gmv_inshop_num,omitempty" xml:"gmv_inshop_num,omitempty"`
    // 拍下订单金额
    GmvInshopAmt   string `json:"gmv_inshop_amt,omitempty" xml:"gmv_inshop_amt,omitempty"`
    // 成交订单数
    AlipayInShopNum   int64 `json:"alipay_in_shop_num,omitempty" xml:"alipay_in_shop_num,omitempty"`
    // 成交订单金额
    AlipayInshopAmt   string `json:"alipay_inshop_amt,omitempty" xml:"alipay_inshop_amt,omitempty"`
    // 展现转化率
    Icvr   string `json:"icvr,omitempty" xml:"icvr,omitempty"`
    // 点击转化率
    Cvr   string `json:"cvr,omitempty" xml:"cvr,omitempty"`
    // 投资回报率
    Roi   string `json:"roi,omitempty" xml:"roi,omitempty"`
    // 记录日期
    LogDate   string `json:"log_date,omitempty" xml:"log_date,omitempty"`
    // 小时
    HourId   int64 `json:"hour_id,omitempty" xml:"hour_id,omitempty"`
    // 计划id
    CampaignId   int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
    // 单元id
    AdgroupId   int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
    // 计划名称
    CampaignName   string `json:"campaign_name,omitempty" xml:"campaign_name,omitempty"`
    // 单元名称
    AdgroupName   string `json:"adgroup_name,omitempty" xml:"adgroup_name,omitempty"`
    // 资源位id
    AdzoneId   int64 `json:"adzone_id,omitempty" xml:"adzone_id,omitempty"`
    // 资源位名称
    AdzoneName   string `json:"adzone_name,omitempty" xml:"adzone_name,omitempty"`
    // 创意id
    CreativeId   int64 `json:"creative_id,omitempty" xml:"creative_id,omitempty"`
    // 创意名称
    CreativeName   string `json:"creative_name,omitempty" xml:"creative_name,omitempty"`
    // 人群id
    CrowdId   int64 `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
    // 人群名称
    CrowdName   string `json:"crowd_name,omitempty" xml:"crowd_name,omitempty"`
}
