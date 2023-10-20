package simba

// TaobaouniversalbplabelfindconfiglistTopResult 结构体
type TaobaouniversalbplabelfindconfiglistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	LabelConfigVOTopBulkData *TopBulkData `json:"label_config_v_o_top_bulk_data,omitempty" xml:"label_config_v_o_top_bulk_data,omitempty"`
}
