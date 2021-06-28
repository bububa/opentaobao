package simba

// CampaignArea 
/* model for simplify = false
type CampaignArea struct {

    // 主人昵称
    
    Nick   string `json:"nick,omitempty"`
    

    // 推广计划ID
    
    CampaignId   int64 `json:"campaign_id,omitempty"`
    

    // 值为：“all”；或者用“,”分割的数字，数字必须是直通车全国省市列表的AreaID，如果已经包含省、自治区id，请不要再包括省内市的id；
    
    Area   string `json:"area,omitempty"`
    

    // 创建时间
    
    CreateTime   string `json:"create_time,omitempty"`
    

    // 最后修改时间
    
    ModifiedTime   string `json:"modified_time,omitempty"`
    

}
*/

// CampaignArea 
type CampaignArea struct {

    // 主人昵称
    Nick   string `json:"nick,omitempty"`

    // 推广计划ID
    CampaignId   int64 `json:"campaign_id,omitempty"`

    // 值为：“all”；或者用“,”分割的数字，数字必须是直通车全国省市列表的AreaID，如果已经包含省、自治区id，请不要再包括省内市的id；
    Area   string `json:"area,omitempty"`

    // 创建时间
    CreateTime   string `json:"create_time,omitempty"`

    // 最后修改时间
    ModifiedTime   string `json:"modified_time,omitempty"`

}
