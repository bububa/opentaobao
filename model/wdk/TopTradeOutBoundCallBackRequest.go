package wdk

// TopTradeOutBoundCallBackRequest 结构体
type TopTradeOutBoundCallBackRequest struct {
	// 子单明细列表
	DemandDetailCallBackRequests []TradeOutBoundDetailCallBackRequest `json:"demand_detail_call_back_requests,omitempty" xml:"demand_detail_call_back_requests>trade_out_bound_detail_call_back_request,omitempty"`
	// 扩展字段
	Extension string `json:"extension,omitempty" xml:"extension,omitempty"`
	// oms主单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 出库状态
	OutBoundStatus string `json:"out_bound_status,omitempty" xml:"out_bound_status,omitempty"`
	// 渠道单号
	ChannelOrderNo string `json:"channel_order_no,omitempty" xml:"channel_order_no,omitempty"`
	// 出库状态发生时间
	DemandStatusTime string `json:"demand_status_time,omitempty" xml:"demand_status_time,omitempty"`
	// 仓编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 渠道类型
	ChannelSource string `json:"channel_source,omitempty" xml:"channel_source,omitempty"`
}
