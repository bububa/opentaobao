package wdk

// CreateReverseResponse 结构体
type CreateReverseResponse struct {
	// 逆向单ids
	ReverseIds []int64 `json:"reverse_ids,omitempty" xml:"reverse_ids>int64,omitempty"`
	// 外部单号
	OutBizOrderIds []string `json:"out_biz_order_ids,omitempty" xml:"out_biz_order_ids>string,omitempty"`
	// tp单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 门店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
