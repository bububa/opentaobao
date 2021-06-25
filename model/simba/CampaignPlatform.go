package simba

// CampaignPlatform 
type CampaignPlatform struct {

    // 主人昵称
    Nick   string `json:"nick,omitempty"`

    // 推广计划ID
    CampaignId   int64 `json:"campaign_id,omitempty"`

    // 已经废弃了
    OutsideDiscount   int64 `json:"outside_discount,omitempty"`

    // 创建时间
    CreateTime   string `json:"create_time,omitempty"`

    // 最后修改时间
    ModifiedTime   string `json:"modified_time,omitempty"`

    // 搜索投放频道代码数组，频道代码必须是直通车搜索类频道列表中的值。1：淘宝站内搜索，7、无线站内搜索；9:无线站外搜索
    SearchChannels   []Number `json:"search_channels,omitempty"`

    // 非搜索投放频道代码数组，频道代码必须是直通车非搜索类频道列表中的值。2、站内定向；4、站外定向；8、无线站内定向；10、无线站外定向
    NonsearchChannels   []Number `json:"nonsearch_channels,omitempty"`

    // 已经废弃了
    MobileDiscount   int64 `json:"mobile_discount,omitempty"`

}