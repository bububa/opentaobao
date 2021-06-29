package simba

// RptBaseEntityDTO 
type RptBaseEntityDTO struct {
    // 日期
    Thedate   string `json:"thedate,omitempty" xml:"thedate,omitempty"`
    // 计划id
    Campaignid   string `json:"campaignid,omitempty" xml:"campaignid,omitempty"`
    // 推广组id
    Adgroupid   string `json:"adgroupid,omitempty" xml:"adgroupid,omitempty"`
    // 花费
    Cost   string `json:"cost,omitempty" xml:"cost,omitempty"`
    // 点击率
    Ctr   string `json:"ctr,omitempty" xml:"ctr,omitempty"`
    // 每点击一次花费
    Cpc   string `json:"cpc,omitempty" xml:"cpc,omitempty"`
    // 每千次展现花费
    Cpm   string `json:"cpm,omitempty" xml:"cpm,omitempty"`
    // 展现量
    Impression   int64 `json:"impression,omitempty" xml:"impression,omitempty"`
    // 点击量
    Click   int64 `json:"click,omitempty" xml:"click,omitempty"`
    // 人群名称
    Crowdname   string `json:"crowdname,omitempty" xml:"crowdname,omitempty"`
    // 流量来源，设备和网络。值包含PC站内，PC站外,移动站内，移动站外
    Source   string `json:"source,omitempty" xml:"source,omitempty"`
    // 人群id
    Crowdid   string `json:"crowdid,omitempty" xml:"crowdid,omitempty"`
}
