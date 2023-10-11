package aliospay

// GetPeriodAgreementStatusRequest 结构体
type GetPeriodAgreementStatusRequest struct {
	// 请求唯一id，不可重复，服务端会根据此参数防重放
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 语言,en表示英文，zh表示中文
	Lang string `json:"lang,omitempty" xml:"lang,omitempty"`
	// 请求时间戳
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 外部商户周期扣款签约码，周期扣款协议中标示用户的唯一签约编号（确保在商户系统中唯一）。格式规则：支持大写小写字母和数字，最长32位。
	ExternalPeriodAgreementCode string `json:"external_period_agreement_code,omitempty" xml:"external_period_agreement_code,omitempty"`
}
