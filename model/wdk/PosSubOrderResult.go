package wdk

// PosSubOrderResult 
type PosSubOrderResult struct {
    // subOrderId
    SubOrderId   int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
    // outOrderId
    OutOrderId   string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}
