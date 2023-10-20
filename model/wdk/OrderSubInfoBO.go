package wdk

// OrderSubInfoBo 结构体
type OrderSubInfoBo struct {
	// 子业务单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 外部子单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}
