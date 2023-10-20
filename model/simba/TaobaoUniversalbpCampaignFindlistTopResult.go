package simba

// TaobaouniversalbpcampaignfindlistTopResult 结构体
type TaobaouniversalbpcampaignfindlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	CampaignVOTopBulkData *TopBulkData `json:"campaign_v_o_top_bulk_data,omitempty" xml:"campaign_v_o_top_bulk_data,omitempty"`
}
