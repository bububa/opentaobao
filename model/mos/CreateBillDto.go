package mos

// CreateBillDto 结构体
type CreateBillDto struct {
	// 结算单列表
	SettlementDTOList []BillSettlementDto `json:"settlement_d_t_o_list,omitempty" xml:"settlement_d_t_o_list>bill_settlement_dto,omitempty"`
	// 收款方国家，国际标准的2位简码
	PayeeCountry string `json:"payee_country,omitempty" xml:"payee_country,omitempty"`
	// 支付币种，国际标准的3位简码
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// //BANK("普通银行账号") ,     //SPECIAL_CN_BANK("专用银行账号"),     //ALIPAY("支付宝"),     //OTHER("第三方支付")
	PayChannels string `json:"pay_channels,omitempty" xml:"pay_channels,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 费用发生公司(门店名)
	ExpenseOrgCode string `json:"expense_org_code,omitempty" xml:"expense_org_code,omitempty"`
	// PREPAY(先付款后收票), POSTPAY(先收票后付款), NONE(无票)
	InvoiceRule string `json:"invoice_rule,omitempty" xml:"invoice_rule,omitempty"`
	// 付款期限，即审批通过后几天付款，不填时以业务模块上的配置为准
	PayTerm int64 `json:"pay_term,omitempty" xml:"pay_term,omitempty"`
	// 费用发生公司(门店名)
	ExpenseOrgId string `json:"expense_org_id,omitempty" xml:"expense_org_id,omitempty"`
	// 业务模块编码
	BizModuleCode string `json:"biz_module_code,omitempty" xml:"biz_module_code,omitempty"`
	// 附件id列表
	AttachmentIds []int64 `json:"attachment_ids,omitempty" xml:"attachment_ids>int64,omitempty"`
	// 实际需求方，代理申请时填写实际的需求方
	ReqmentApplicant string `json:"reqment_applicant,omitempty" xml:"reqment_applicant,omitempty"`
	// //付款单总金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 是否自动提交审批
	AutoSumbit bool `json:"auto_sumbit,omitempty" xml:"auto_sumbit,omitempty"`
	// 单据请求唯一编号，必填
	BillNo string `json:"bill_no,omitempty" xml:"bill_no,omitempty"`
	// 财务初审，如果是免审的单据，填财务工号
	FinancePreWorkNo string `json:"finance_pre_work_no,omitempty" xml:"finance_pre_work_no,omitempty"`
	// 申请人工号
	Applicant string `json:"applicant,omitempty" xml:"applicant,omitempty"`
	// 审批通过时间，如果是免审的单据，该字段必填
	ApproveDate string `json:"approve_date,omitempty" xml:"approve_date,omitempty"`
	// 扩展
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// 付款说明，该说明会带到智付的订单上，做为订单说明
	Comments string `json:"comments,omitempty" xml:"comments,omitempty"`
}
