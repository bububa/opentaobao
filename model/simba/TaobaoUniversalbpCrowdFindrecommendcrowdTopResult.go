package simba

// TaobaoUniversalbpCrowdFindrecommendcrowdTopResult 结构体
type TaobaoUniversalbpCrowdFindrecommendcrowdTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	CrowdRefVOTopBulkData *TopBulkData `json:"crowd_ref_v_o_top_bulk_data,omitempty" xml:"crowd_ref_v_o_top_bulk_data,omitempty"`
}
