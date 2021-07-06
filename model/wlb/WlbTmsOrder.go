package wlb

// WlbTmsOrder 结构体
type WlbTmsOrder struct {
	// 物流订单编号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 运单号
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 物流公司编码
	CompanyCode string `json:"company_code,omitempty" xml:"company_code,omitempty"`
	// 物流订单的所有者ID,货主
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
