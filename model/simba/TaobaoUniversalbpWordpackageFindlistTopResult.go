package simba

// TaobaouniversalbpwordpackagefindlistTopResult 结构体
type TaobaouniversalbpwordpackagefindlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	WordPackageVOTopBulkData *TopBulkData `json:"word_package_v_o_top_bulk_data,omitempty" xml:"word_package_v_o_top_bulk_data,omitempty"`
}
