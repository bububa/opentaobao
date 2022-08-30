package ascp

// DeliveryOrderConfirmReportRequest 结构体
type DeliveryOrderConfirmReportRequest struct {
	// 包裹信息
	Packages []DeliveryOrder `json:"packages,omitempty" xml:"packages>delivery_order,omitempty"`
	// 单据列表
	OrderLines []OrderLine `json:"order_lines,omitempty" xml:"order_lines>order_line,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 扩展属性
	ExtendProps string `json:"extend_props,omitempty" xml:"extend_props,omitempty"`
	// 货主ID
	OwnerCode string `json:"owner_code,omitempty" xml:"owner_code,omitempty"`
	// 发货单信息
	DeliveryOrder *DeliveryOrder `json:"delivery_order,omitempty" xml:"delivery_order,omitempty"`
	// 业务请求时间（时间戳）
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}
