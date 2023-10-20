package simba

// TaobaoUniversalbpWordpackageSuggestkrlistTopResult 结构体
type TaobaoUniversalbpWordpackageSuggestkrlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	SuggestWordPackageVOTopBulkData *TopBulkData `json:"suggest_word_package_v_o_top_bulk_data,omitempty" xml:"suggest_word_package_v_o_top_bulk_data,omitempty"`
}
