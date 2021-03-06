package tbk

// PageResult 结构体
type PageResult struct {
	// 处罚订单列表
	Results []TaobaoTbkDgPunishOrderGetResult `json:"results,omitempty" xml:"results>taobao_tbk_dg_punish_order_get_result,omitempty"`
	// 翻页的pageno
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 翻页的pagesie
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 一共能查询出来的结果总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
