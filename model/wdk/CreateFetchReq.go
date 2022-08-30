package wdk

// CreateFetchReq 结构体
type CreateFetchReq struct {
	// 取货详情
	FetchAggregateList []FetchAggregate `json:"fetch_aggregate_list,omitempty" xml:"fetch_aggregate_list>fetch_aggregate,omitempty"`
	// 买家地址
	BuyerAddress string `json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
	// 买家名称
	BuyerName string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// 买家电话
	BuyerPhone string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
	// 期望取货结束时间
	ExpectFetchEndTime string `json:"expect_fetch_end_time,omitempty" xml:"expect_fetch_end_time,omitempty"`
	// 期望取货开始时间
	ExpectFetchStartTime string `json:"expect_fetch_start_time,omitempty" xml:"expect_fetch_start_time,omitempty"`
	// 主单号
	MainOutOrderId string `json:"main_out_order_id,omitempty" xml:"main_out_order_id,omitempty"`
	// 备注
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 退款单号 幂等id
	ReverseId string `json:"reverse_id,omitempty" xml:"reverse_id,omitempty"`
	// 门店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 上门取货（1为上门取货）
	FetchType int64 `json:"fetch_type,omitempty" xml:"fetch_type,omitempty"`
	// 发起人
	Operator *OperatorVo `json:"operator,omitempty" xml:"operator,omitempty"`
	// 原因id
	ReasonId int64 `json:"reason_id,omitempty" xml:"reason_id,omitempty"`
	// 退款金额
	RefundAmount int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
}
