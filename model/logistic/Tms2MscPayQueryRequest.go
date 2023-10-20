package logistic

// Tms2mscPayQueryRequest 结构体
type Tms2mscPayQueryRequest struct {
	// 服务类型
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 配资源编码
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 唯一标识单号
	DeliveryCode string `json:"delivery_code,omitempty" xml:"delivery_code,omitempty"`
}
