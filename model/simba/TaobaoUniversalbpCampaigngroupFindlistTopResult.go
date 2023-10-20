package simba

// TaobaoUniversalbpCampaigngroupFindlistTopResult 结构体
type TaobaoUniversalbpCampaigngroupFindlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	CampaignGroupVOTopBulkData *TopBulkData `json:"campaign_group_v_o_top_bulk_data,omitempty" xml:"campaign_group_v_o_top_bulk_data,omitempty"`
}
