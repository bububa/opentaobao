package omniorder

// Door2doorQueryResult 结构体
type Door2doorQueryResult struct {
	// 码对应的淘宝主订单ID
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
}
