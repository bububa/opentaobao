package examination

// ReportImTokenStatusRequest 结构体
type ReportImTokenStatusRequest struct {
	// 令牌值
	ImToken string `json:"im_token,omitempty" xml:"im_token,omitempty"`
	// 订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
