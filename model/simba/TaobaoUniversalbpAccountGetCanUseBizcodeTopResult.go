package simba

// TaobaoUniversalbpAccountGetCanUseBizcodeTopResult 结构体
type TaobaoUniversalbpAccountGetCanUseBizcodeTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	TopMarketSceneVOTopBulkData *TopBulkData `json:"top_market_scene_v_o_top_bulk_data,omitempty" xml:"top_market_scene_v_o_top_bulk_data,omitempty"`
}
