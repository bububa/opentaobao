package wdk

// PosSubOrderResult 结构体
type PosSubOrderResult struct {
	// outOrderId
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// subOrderId
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
}
