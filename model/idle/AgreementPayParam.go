package idle

// AgreementPayParam 结构体
type AgreementPayParam struct {
	// 商户订单号（唯一建）
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 代扣计划
	PlanId string `json:"plan_id,omitempty" xml:"plan_id,omitempty"`
	// R1:闲鱼回收代扣（V1版本）
	PayBizCode string `json:"pay_biz_code,omitempty" xml:"pay_biz_code,omitempty"`
}
