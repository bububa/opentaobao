package simba

// TaobaoUniversalbpMaterialItemFindpageTopResult 结构体
type TaobaoUniversalbpMaterialItemFindpageTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	ItemVOTopBulkData *TopBulkData `json:"item_v_o_top_bulk_data,omitempty" xml:"item_v_o_top_bulk_data,omitempty"`
}
