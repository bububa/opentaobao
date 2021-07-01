package tmallservice

// Paged 结构体
type Paged struct {
	// 工单列表
	DataList []SpServiceOrderDto `json:"data_list,omitempty" xml:"data_list>sp_service_order_dto,omitempty"`
	// 总页数
	TotalPageCount int64 `json:"total_page_count,omitempty" xml:"total_page_count,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总记录数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
