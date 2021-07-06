package wdk

// FinanceOrderDetailResponse 结构体
type FinanceOrderDetailResponse struct {
	// 财务订单信息
	FinanceOrderDetails []FinanceOrderDetail `json:"finance_order_details,omitempty" xml:"finance_order_details>finance_order_detail,omitempty"`
	// 分页信息
	Pagination *Pagination `json:"pagination,omitempty" xml:"pagination,omitempty"`
}
