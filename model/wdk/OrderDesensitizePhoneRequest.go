package wdk

// OrderDesensitizePhoneRequest 结构体
type OrderDesensitizePhoneRequest struct {
	// 业务单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 经营店编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
}
