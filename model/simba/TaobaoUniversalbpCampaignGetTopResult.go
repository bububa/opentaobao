package simba

// TaobaouniversalbpcampaigngetTopResult 结构体
type TaobaouniversalbpcampaigngetTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	TopCampaignVO *TopCampaignVo `json:"top_campaign_v_o,omitempty" xml:"top_campaign_v_o,omitempty"`
}
