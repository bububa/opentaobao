package simba

// TaobaouniversalbpadzonefindconfiglistTopResult 结构体
type TaobaouniversalbpadzonefindconfiglistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	AdzoneConfigVOTopBulkData *TopBulkData `json:"adzone_config_v_o_top_bulk_data,omitempty" xml:"adzone_config_v_o_top_bulk_data,omitempty"`
}
