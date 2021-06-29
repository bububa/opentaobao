package simba

// CampaignSchedule 
type CampaignSchedule struct {
    // 主人昵称
    Nick   string `json:"nick,omitempty" xml:"nick,omitempty"`
    // 推广计划ID
    CampaignId   int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
    // 值为：“all”；或者用“;”分割的每天的设置字符串，该字符串为用“,”分割的时段折扣字符串，格式为：起始时间-结束时间:折扣，其中时间是24小时格式记录18:30，，折扣是1-150整数，表示折扣百分比；
    Schedule   string `json:"schedule,omitempty" xml:"schedule,omitempty"`
    // 创建时间
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`
    // 最后修改时间
    ModifiedTime   string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
}
