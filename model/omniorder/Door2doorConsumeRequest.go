package omniorder

// Door2doorConsumeRequest 结构体
type Door2doorConsumeRequest struct {
	// 核销码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 淘宝主订单ID
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
}
