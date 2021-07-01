package promotion

// Page 结构体
type Page struct {
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 每页查询数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 结果
	Datas []AlibabaAsrDataservicePromotionruleQueryData `json:"datas,omitempty" xml:"datas>alibaba_asr_dataservice_promotionrule_query_data,omitempty"`
	// 活动列表
	Activities []ActivityDto `json:"activities,omitempty" xml:"activities>activity_dto,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 总记录数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
