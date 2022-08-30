package wdk

// LoadFetchResponse 结构体
type LoadFetchResponse struct {
	// 取货单list
	FetchAggregates []FetchAggregateSdo `json:"fetch_aggregates,omitempty" xml:"fetch_aggregates>fetch_aggregate_sdo,omitempty"`
	// 退货取货单ID
	FetchOrderId string `json:"fetch_order_id,omitempty" xml:"fetch_order_id,omitempty"`
}
