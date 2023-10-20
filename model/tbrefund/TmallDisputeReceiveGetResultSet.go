package tbrefund

// TmalldisputereceivegetResultSet 结构体
type TmalldisputereceivegetResultSet struct {
	// results
	Results []Dispute `json:"results,omitempty" xml:"results>dispute,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 当前页面的纠纷单数量
	PageResults int64 `json:"page_results,omitempty" xml:"page_results,omitempty"`
	// 所有符合查询条件的纠纷单数量
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否还有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}
