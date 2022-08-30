package wdk

// CancelRequest 结构体
type CancelRequest struct {
	// 子订单号列表
	SubBizOrderCodes []string `json:"sub_biz_order_codes,omitempty" xml:"sub_biz_order_codes>string,omitempty"`
	// 订单号
	SourceOrderCode string `json:"source_order_code,omitempty" xml:"source_order_code,omitempty"`
	// 仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 渠道来源
	SourceFrom string `json:"source_from,omitempty" xml:"source_from,omitempty"`
	// 出库单单据类型
	OutboundOrderType int64 `json:"outbound_order_type,omitempty" xml:"outbound_order_type,omitempty"`
}
