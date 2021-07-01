package idle

// AgreementPayBillQueryParam 结构体
type AgreementPayBillQueryParam struct {
	// 订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 业务编码：R1:回收
	PayBizCode string `json:"pay_biz_code,omitempty" xml:"pay_biz_code,omitempty"`
}
