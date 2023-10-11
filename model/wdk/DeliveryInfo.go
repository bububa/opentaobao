package wdk

// DeliveryInfo 结构体
type DeliveryInfo struct {
	// 送货人名称
	DeliveryName string `json:"delivery_name,omitempty" xml:"delivery_name,omitempty"`
	// 送货人手机号
	DeliveryPhone string `json:"delivery_phone,omitempty" xml:"delivery_phone,omitempty"`
}
