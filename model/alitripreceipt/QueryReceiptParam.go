package alitripreceipt

// QueryReceiptParam 结构体
type QueryReceiptParam struct {
	// 扩展参数
	ExtMap string `json:"ext_map,omitempty" xml:"ext_map,omitempty"`
	// 发票流水号
	ReceiptNumber string `json:"receipt_number,omitempty" xml:"receipt_number,omitempty"`
	// 商家id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 分页索引，从0开始
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 订单号
	TpOrderId int64 `json:"tp_order_id,omitempty" xml:"tp_order_id,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 用户id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
