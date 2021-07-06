package tmallhk

// AwdcInfo 结构体
type AwdcInfo struct {
	// 货品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 订单号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 商品ID
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 子订单号
	SubOrderNo string `json:"sub_order_no,omitempty" xml:"sub_order_no,omitempty"`
	// 溯源码
	TraceCode string `json:"trace_code,omitempty" xml:"trace_code,omitempty"`
	// 溯源码镭射时间
	TraceCodeTime string `json:"trace_code_time,omitempty" xml:"trace_code_time,omitempty"`
	// hrd info
	Hrd *AwdcHrd `json:"hrd,omitempty" xml:"hrd,omitempty"`
	// ngtc info
	Ngtc *AwdcNgtc `json:"ngtc,omitempty" xml:"ngtc,omitempty"`
	// shipment
	Shipment *AwdcShipment `json:"shipment,omitempty" xml:"shipment,omitempty"`
}
