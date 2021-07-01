package trade

// CainiaoCntecSupplierOrderServiceResult 结构体
type CainiaoCntecSupplierOrderServiceResult struct {
	// 是否有下一页
	HasNextPage bool `json:"has_next_page,omitempty" xml:"has_next_page,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 分页游标
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 订单列表
	OrderList []SupplierOrder `json:"order_list,omitempty" xml:"order_list>supplier_order,omitempty"`
}
