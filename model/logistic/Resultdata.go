package logistic

// Resultdata 结构体
type Resultdata struct {
	// 主交易单号
	ParentOrderId int64 `json:"parent_order_id,omitempty" xml:"parent_order_id,omitempty"`
}
