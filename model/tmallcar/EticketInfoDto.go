package tmallcar

// EticketInfoDto 结构体
type EticketInfoDto struct {
	// 核销状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 核销状态描述
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
