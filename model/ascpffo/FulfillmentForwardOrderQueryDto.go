package ascpffo

// FulfillmentForwardOrderQueryDto 结构体
type FulfillmentForwardOrderQueryDto struct {
	// 用户订单Id列表
	CustomerOrderNumberList []string `json:"customer_order_number_list,omitempty" xml:"customer_order_number_list>string,omitempty"`
	// 发货单号
	FulfillmentOrderNo string `json:"fulfillment_order_no,omitempty" xml:"fulfillment_order_no,omitempty"`
	// 账套编码
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 分页页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
