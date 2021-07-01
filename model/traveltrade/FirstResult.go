package traveltrade

// FirstResult 结构体
type FirstResult struct {
	// 预约订单ID
	BookOrderId int64 `json:"book_order_id,omitempty" xml:"book_order_id,omitempty"`
	// TC子订单号
	SubTcOrderId int64 `json:"sub_tc_order_id,omitempty" xml:"sub_tc_order_id,omitempty"`
	// 预约ID(主键)
	BookInfoId int64 `json:"book_info_id,omitempty" xml:"book_info_id,omitempty"`
	// TC主订单号
	TcOrderId int64 `json:"tc_order_id,omitempty" xml:"tc_order_id,omitempty"`
}
