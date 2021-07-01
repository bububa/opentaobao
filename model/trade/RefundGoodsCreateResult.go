package trade

// RefundGoodsCreateResult 结构体
type RefundGoodsCreateResult struct {
	// 退货单id
	RefundGoodsId string `json:"refund_goods_id,omitempty" xml:"refund_goods_id,omitempty"`
	// 是否创建成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
