package alitripmerchant

// DerbyVoucherCardCancelDto 结构体
type DerbyVoucherCardCancelDto struct {
	// 订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 订单状态code
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 模拟token
	FkToken string `json:"fk_token,omitempty" xml:"fk_token,omitempty"`
}
