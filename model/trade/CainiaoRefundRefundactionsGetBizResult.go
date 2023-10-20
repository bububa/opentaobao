package trade

// CainiaoRefundRefundactionsGetBizResult 结构体
type CainiaoRefundRefundactionsGetBizResult struct {
	// 调用错误时，错误code
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 调用错误时，错误描述
	StatusMessage string `json:"status_message,omitempty" xml:"status_message,omitempty"`
	// 该订单支持的退款退货操作的集合
	Data *OrderRefundActionResponse `json:"data,omitempty" xml:"data,omitempty"`
	// true表示成功，false表示失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
