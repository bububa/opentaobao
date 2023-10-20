package simba

// TaobaouniversalbpbidwordfindlistTopResult 结构体
type TaobaouniversalbpbidwordfindlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	WordVOTopBulkData *TopBulkData `json:"word_v_o_top_bulk_data,omitempty" xml:"word_v_o_top_bulk_data,omitempty"`
}
