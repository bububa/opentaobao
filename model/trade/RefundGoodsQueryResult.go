package trade

// RefundGoodsQueryResult 结构体
type RefundGoodsQueryResult struct {
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 取货类型（"FETCH_HOME"：上门；"ON_SHOP"：到店；"NONE"：无需取）
	RefundFetchType string `json:"refund_fetch_type,omitempty" xml:"refund_fetch_type,omitempty"`
	// 买家id
	BuyerId string `json:"buyer_id,omitempty" xml:"buyer_id,omitempty"`
	// 买家姓名
	BuyerName string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// 买家电话
	BuyerPhone string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
	// 买家地址
	BuyerAddress string `json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
	// 发起退货来源
	InitFrom int64 `json:"init_from,omitempty" xml:"init_from,omitempty"`
	// 发起人
	InitOperator string `json:"init_operator,omitempty" xml:"init_operator,omitempty"`
	// 发起退货备注
	InitMemo string `json:"init_memo,omitempty" xml:"init_memo,omitempty"`
	// 退货子订单详情
	RefundGoodsSubOrderDetailList []RefundGoodsSubOrderDetail `json:"refund_goods_sub_order_detail_list,omitempty" xml:"refund_goods_sub_order_detail_list>refund_goods_sub_order_detail,omitempty"`
}
