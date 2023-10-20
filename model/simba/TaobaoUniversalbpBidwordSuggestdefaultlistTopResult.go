package simba

// TaobaouniversalbpbidwordsuggestdefaultlistTopResult 结构体
type TaobaouniversalbpbidwordsuggestdefaultlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	BidwordSuggestItemVOTopBulkData *TopBulkData `json:"bidword_suggest_item_v_o_top_bulk_data,omitempty" xml:"bidword_suggest_item_v_o_top_bulk_data,omitempty"`
}
