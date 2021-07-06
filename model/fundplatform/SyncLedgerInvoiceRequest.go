package fundplatform

// SyncLedgerInvoiceRequest 结构体
type SyncLedgerInvoiceRequest struct {
	// 发票行
	InputInvoiceLineDTOList []InputInvoiceLineLedgerDto `json:"input_invoice_line_d_t_o_list,omitempty" xml:"input_invoice_line_d_t_o_list>input_invoice_line_ledger_dto,omitempty"`
	// 金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 认证时间
	CertifyDate string `json:"certify_date,omitempty" xml:"certify_date,omitempty"`
	// 认证标识
	CertifyFlag string `json:"certify_flag,omitempty" xml:"certify_flag,omitempty"`
	// 校验码
	CheckSum string `json:"check_sum,omitempty" xml:"check_sum,omitempty"`
	// 密码区
	CipherText string `json:"cipher_text,omitempty" xml:"cipher_text,omitempty"`
	// 不含税金额
	ExcludingTaxAmount string `json:"excluding_tax_amount,omitempty" xml:"excluding_tax_amount,omitempty"`
	// 发票代码
	InvoiceCode string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty"`
	// 开票日期
	InvoiceDate string `json:"invoice_date,omitempty" xml:"invoice_date,omitempty"`
	// 发票号码
	InvoiceNo string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
	// 发票状态
	InvoiceStatus string `json:"invoice_status,omitempty" xml:"invoice_status,omitempty"`
	// 发票类型
	InvoiceType string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 机器编码
	MachineCode string `json:"machine_code,omitempty" xml:"machine_code,omitempty"`
	// 购方银行信息
	PurchaserBankInfo string `json:"purchaser_bank_info,omitempty" xml:"purchaser_bank_info,omitempty"`
	// 购方联系方式
	PurchaserContactInfo string `json:"purchaser_contact_info,omitempty" xml:"purchaser_contact_info,omitempty"`
	// 购方名称
	PurchaserName string `json:"purchaser_name,omitempty" xml:"purchaser_name,omitempty"`
	// 购方税号
	PurchaserTaxNo string `json:"purchaser_tax_no,omitempty" xml:"purchaser_tax_no,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 销方银行信息
	SellerBankInfo string `json:"seller_bank_info,omitempty" xml:"seller_bank_info,omitempty"`
	// 销方联系人信息
	SellerContactInfo string `json:"seller_contact_info,omitempty" xml:"seller_contact_info,omitempty"`
	// 销方名称
	SellerName string `json:"seller_name,omitempty" xml:"seller_name,omitempty"`
	// 销方税号
	SellerTaxNo string `json:"seller_tax_no,omitempty" xml:"seller_tax_no,omitempty"`
	// 税额
	TaxAmount string `json:"tax_amount,omitempty" xml:"tax_amount,omitempty"`
	// 所属期
	TaxPeriod string `json:"tax_period,omitempty" xml:"tax_period,omitempty"`
	// 勾选时间
	CheckCertifyTime string `json:"check_certify_time,omitempty" xml:"check_certify_time,omitempty"`
	// 勾选状态
	CheckCertifyStatus string `json:"check_certify_status,omitempty" xml:"check_certify_status,omitempty"`
	// 有效税额
	EffectiveTaxAmount string `json:"effective_tax_amount,omitempty" xml:"effective_tax_amount,omitempty"`
	// 认证方式标识
	TaxDeductFag string `json:"tax_deduct_fag,omitempty" xml:"tax_deduct_fag,omitempty"`
}
