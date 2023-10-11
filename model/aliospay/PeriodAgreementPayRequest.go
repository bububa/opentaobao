package aliospay

// PeriodAgreementPayRequest 结构体
type PeriodAgreementPayRequest struct {
	// 请求唯一id，不可重复，服务端会根据此参数防重放
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 语言,en表示英文，zh表示中文
	Lang string `json:"lang,omitempty" xml:"lang,omitempty"`
	// 请求时间戳
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 业务订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 订单标题
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 外部商户周期扣款签约码，周期扣款协议中标示用户的唯一签约编号（确保在商户系统中唯一）。格式规则：支持大写小写字母和数字，最长32位
	ExternalPeriodAgreementCode string `json:"external_period_agreement_code,omitempty" xml:"external_period_agreement_code,omitempty"`
	// cp服务端支持的协议，目前只支持HTTPS
	ServiceProtocol string `json:"service_protocol,omitempty" xml:"service_protocol,omitempty"`
	// 支付结果回调通知URL
	PayNotifyUrl string `json:"pay_notify_url,omitempty" xml:"pay_notify_url,omitempty"`
	// 订单总金额，单位为分
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
}
