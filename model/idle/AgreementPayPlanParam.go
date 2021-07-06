package idle

// AgreementPayPlanParam 结构体
type AgreementPayPlanParam struct {
	// 需要拆分的扣款计划,注意只能拆一次只能支持5个 例如[10000,2300]
	Plans []int64 `json:"plans,omitempty" xml:"plans>int64,omitempty"`
	// 商户订单号（唯一建）
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// R1:闲鱼回收代扣（V1版本）
	PayBizCode string `json:"pay_biz_code,omitempty" xml:"pay_biz_code,omitempty"`
	// 订单代扣总金额（单位分）
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
}
