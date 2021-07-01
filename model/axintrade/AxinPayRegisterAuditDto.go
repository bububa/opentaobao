package axintrade

// AxinPayRegisterAuditDto 结构体
type AxinPayRegisterAuditDto struct {
	// 审核结果
	AuditResult bool `json:"audit_result,omitempty" xml:"audit_result,omitempty"`
	// 商户code
	ExternalId string `json:"external_id,omitempty" xml:"external_id,omitempty"`
	// 备注
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 支付平台入驻申请单号
	PayPlatformApplyOrder string `json:"pay_platform_apply_order,omitempty" xml:"pay_platform_apply_order,omitempty"`
	// 支付入驻域申请单号
	PayRegisterOrderNo string `json:"pay_register_order_no,omitempty" xml:"pay_register_order_no,omitempty"`
	// 拒绝原因code
	RejectReasonCode string `json:"reject_reason_code,omitempty" xml:"reject_reason_code,omitempty"`
	// 拒绝原因描述
	RejectReasonDesc string `json:"reject_reason_desc,omitempty" xml:"reject_reason_desc,omitempty"`
	// 支付平台商家编码,smid
	PayMerchantCode string `json:"pay_merchant_code,omitempty" xml:"pay_merchant_code,omitempty"`
}
