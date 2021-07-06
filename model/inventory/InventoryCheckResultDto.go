package inventory

// InventoryCheckResultDto 结构体
type InventoryCheckResultDto struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 每个货品的调整子单据号，作为业务调整依据，处理时会幂等
	SubOrderId string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 每个子调整单据是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
