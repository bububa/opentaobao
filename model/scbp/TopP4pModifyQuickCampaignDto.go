package scbp

// TopP4pModifyQuickCampaignDTO 
type TopP4pModifyQuickCampaignDTO struct {
    // 操作类型，0=计划暂停，1=计划开启，2=计划删除
    Action   int64 `json:"action,omitempty" xml:"action,omitempty"`
    // 计划ID
    CampaignId   int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}
