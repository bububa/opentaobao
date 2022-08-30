package logistic

// LabelDto 结构体
type LabelDto struct {
	// 物流订单号
	LogisticsOrderId string `json:"logistics_order_id,omitempty" xml:"logistics_order_id,omitempty"`
	// 面单链接
	LabelUrl string `json:"label_url,omitempty" xml:"label_url,omitempty"`
}
