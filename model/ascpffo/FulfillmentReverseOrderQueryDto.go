package ascpffo

// FulfillmentReverseOrderQueryDto 结构体
type FulfillmentReverseOrderQueryDto struct {
	// 账套编码
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 用户订单号列表
	CustomerOrderNumberList []string `json:"customer_order_number_list,omitempty" xml:"customer_order_number_list>string,omitempty"`
	// 履约单号
	FulfillmentOrderNo string `json:"fulfillment_order_no,omitempty" xml:"fulfillment_order_no,omitempty"`
	// 分页页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
