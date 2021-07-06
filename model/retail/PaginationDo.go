package retail

// PaginationDo 结构体
type PaginationDo struct {
	// 数据
	DataList []VendingBizOrderDto `json:"data_list,omitempty" xml:"data_list>vending_biz_order_dto,omitempty"`
	// 总数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}
