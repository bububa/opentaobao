package tmallcar

// CreditReceiveDto 结构体
type CreditReceiveDto struct {
	// 扩展字段2
	ExtensionField02 string `json:"extension_field02,omitempty" xml:"extension_field02,omitempty"`
	// 扩展字段1
	ExtensionField01 string `json:"extension_field01,omitempty" xml:"extension_field01,omitempty"`
	// 金融服务商名称
	CapitalName string `json:"capital_name,omitempty" xml:"capital_name,omitempty"`
	// 金融服务商编码
	CapitalCode string `json:"capital_code,omitempty" xml:"capital_code,omitempty"`
	// 授信未通过原因描述
	FailReason string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	// 如果授信通过了,要有授信额度(单位分)
	CreditAmount int64 `json:"credit_amount,omitempty" xml:"credit_amount,omitempty"`
	// 金融机构的授信申请内部单号
	CreditId string `json:"credit_id,omitempty" xml:"credit_id,omitempty"`
	// 授信结果
	CreditResult string `json:"credit_result,omitempty" xml:"credit_result,omitempty"`
	// 天猫业务单号
	TmallBizNo string `json:"tmall_biz_no,omitempty" xml:"tmall_biz_no,omitempty"`
	// 业务类型
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 请求时间
	RequestTime string `json:"request_time,omitempty" xml:"request_time,omitempty"`
}
