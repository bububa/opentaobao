package ascpchannel

// Detailorder 结构体
type Detailorder struct {
	// 实际操作子单id(例如:ICP子单,,UDP子单)
	OperationDetailOrderId string `json:"operation_detail_order_id,omitempty" xml:"operation_detail_order_id,omitempty"`
}
