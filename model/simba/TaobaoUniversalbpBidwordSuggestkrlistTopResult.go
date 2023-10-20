package simba

// TaobaouniversalbpbidwordsuggestkrlistTopResult 结构体
type TaobaouniversalbpbidwordsuggestkrlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	SuggestBidwordVOTopBulkData *TopBulkData `json:"suggest_bidword_v_o_top_bulk_data,omitempty" xml:"suggest_bidword_v_o_top_bulk_data,omitempty"`
}
