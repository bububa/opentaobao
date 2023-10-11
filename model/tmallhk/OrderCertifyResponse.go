package tmallhk

// OrderCertifyResponse 结构体
type OrderCertifyResponse struct {
	// 具体实名信息
	OrderCertify *OrderCertify `json:"order_certify,omitempty" xml:"order_certify,omitempty"`
	// 是否已经实名
	Auth bool `json:"auth,omitempty" xml:"auth,omitempty"`
}
