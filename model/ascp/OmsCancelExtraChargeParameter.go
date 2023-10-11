package ascp

// OmsCancelExtraChargeParameter 结构体
type OmsCancelExtraChargeParameter struct {
	// BFC单号
	WdsCoordinationOrderId string `json:"wds_coordination_order_id,omitempty" xml:"wds_coordination_order_id,omitempty"`
	// 增加费用服务调整单ID
	ExtraChargeServiceOrderId string `json:"extra_charge_service_order_id,omitempty" xml:"extra_charge_service_order_id,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 扩展字段
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 商家id
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}
