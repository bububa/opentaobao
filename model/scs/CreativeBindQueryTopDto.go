package scs

// CreativeBindQueryTopDto 结构体
type CreativeBindQueryTopDto struct {
	// campaignIdList
	CampaignIdList []int64 `json:"campaign_id_list,omitempty" xml:"campaign_id_list>int64,omitempty"`
	// status
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// bizCode
	BizCod string `json:"biz_cod,omitempty" xml:"biz_cod,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// adgroupId
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 报表查询参数
	ReportQuery *ReportQueryTopDto `json:"report_query,omitempty" xml:"report_query,omitempty"`
}
