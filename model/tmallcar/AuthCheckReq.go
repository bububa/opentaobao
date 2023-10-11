package tmallcar

// AuthCheckReq 结构体
type AuthCheckReq struct {
	// 鉴权token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
