package tblogistics

// ConsignDto 结构体
type ConsignDto struct {
	// 发货文案描述
	ConsignDesc string `json:"consign_desc,omitempty" xml:"consign_desc,omitempty"`
	// 物流发货单号
	LpOrderId int64 `json:"lp_order_id,omitempty" xml:"lp_order_id,omitempty"`
}
