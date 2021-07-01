package qimen

// TaobaoQimenStockoutConfirmStruct 结构体
type TaobaoQimenStockoutConfirmStruct struct {
	// deliveryOrder
	DeliveryOrder *DeliveryOrder `json:"deliveryOrder,omitempty" xml:"deliveryOrder,omitempty"`
	// packages
	Packages []Package `json:"packages,omitempty" xml:"packages>package,omitempty"`
	// orderLines
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenStockoutConfirmMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
