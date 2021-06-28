package simba

// RptBaseEntityDTO 
/* model for simplify = false
type RptBaseEntityDTO struct {

    // 日期
    
    Thedate   string `json:"thedate,omitempty"`
    

    // 计划id
    
    Campaignid   string `json:"campaignid,omitempty"`
    

    // 推广组id
    
    Adgroupid   string `json:"adgroupid,omitempty"`
    

    // 花费
    
    Cost   string `json:"cost,omitempty"`
    

    // 点击率
    
    Ctr   string `json:"ctr,omitempty"`
    

    // 每点击一次花费
    
    Cpc   string `json:"cpc,omitempty"`
    

    // 每千次展现花费
    
    Cpm   string `json:"cpm,omitempty"`
    

    // 展现量
    
    Impression   int64 `json:"impression,omitempty"`
    

    // 点击量
    
    Click   int64 `json:"click,omitempty"`
    

    // 人群名称
    
    Crowdname   string `json:"crowdname,omitempty"`
    

    // 流量来源，设备和网络。值包含PC站内，PC站外,移动站内，移动站外
    
    Source   string `json:"source,omitempty"`
    

    // 人群id
    
    Crowdid   string `json:"crowdid,omitempty"`
    

}
*/

// RptBaseEntityDTO 
type RptBaseEntityDTO struct {

    // 日期
    Thedate   string `json:"thedate,omitempty"`

    // 计划id
    Campaignid   string `json:"campaignid,omitempty"`

    // 推广组id
    Adgroupid   string `json:"adgroupid,omitempty"`

    // 花费
    Cost   string `json:"cost,omitempty"`

    // 点击率
    Ctr   string `json:"ctr,omitempty"`

    // 每点击一次花费
    Cpc   string `json:"cpc,omitempty"`

    // 每千次展现花费
    Cpm   string `json:"cpm,omitempty"`

    // 展现量
    Impression   int64 `json:"impression,omitempty"`

    // 点击量
    Click   int64 `json:"click,omitempty"`

    // 人群名称
    Crowdname   string `json:"crowdname,omitempty"`

    // 流量来源，设备和网络。值包含PC站内，PC站外,移动站内，移动站外
    Source   string `json:"source,omitempty"`

    // 人群id
    Crowdid   string `json:"crowdid,omitempty"`

}
