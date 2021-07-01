package lstlogistics

// SubOrderParam 结构体
type SubOrderParam struct {
	// 子订单id
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 购买数量
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}
