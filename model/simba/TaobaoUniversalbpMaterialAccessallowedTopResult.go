package simba

// TaobaoUniversalbpMaterialAccessallowedTopResult 结构体
type TaobaoUniversalbpMaterialAccessallowedTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	MaterialAccessAllowVOTopBulkData *TopBulkData `json:"material_access_allow_v_o_top_bulk_data,omitempty" xml:"material_access_allow_v_o_top_bulk_data,omitempty"`
}
