package logistic

// AeopWlCreateWarehouseOrderResultDto 结构体
type AeopWlCreateWarehouseOrderResultDto struct {
	// 创建时错误信息
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// 国际运单号
	IntlTrackingNo string `json:"intl_tracking_no,omitempty" xml:"intl_tracking_no,omitempty"`
	// 订单来源
	TradeOrderFrom string `json:"trade_order_from,omitempty" xml:"trade_order_from,omitempty"`
	// 创建时错误码(1表示无错误)
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 外部订单号
	OutOrderId int64 `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 交易订单号
	TradeOrderId int64 `json:"trade_order_id,omitempty" xml:"trade_order_id,omitempty"`
	// 物流订单号
	WarehouseOrderId int64 `json:"warehouse_order_id,omitempty" xml:"warehouse_order_id,omitempty"`
	// 创建订单是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
