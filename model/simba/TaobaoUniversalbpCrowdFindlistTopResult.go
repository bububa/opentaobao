package simba

// TaobaoUniversalbpCrowdFindlistTopResult 结构体
type TaobaoUniversalbpCrowdFindlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	CrowdBindResultVOTopBulkData *TopBulkData `json:"crowd_bind_result_v_o_top_bulk_data,omitempty" xml:"crowd_bind_result_v_o_top_bulk_data,omitempty"`
}
