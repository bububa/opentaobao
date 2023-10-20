package simba

// TaobaoUniversalbpCreativeManageFindmanagepageTopResult 结构体
type TaobaoUniversalbpCreativeManageFindmanagepageTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	CreativeVOTopBulkData *TopBulkData `json:"creative_v_o_top_bulk_data,omitempty" xml:"creative_v_o_top_bulk_data,omitempty"`
}
