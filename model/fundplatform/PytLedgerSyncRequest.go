package fundplatform

import (
	"sync"
)

// PytLedgerSyncRequest 结构体
type PytLedgerSyncRequest struct {
	// 发票行信息
	Details []PytLedgerInvoiceLineRequest `json:"details,omitempty" xml:"details>pyt_ledger_invoice_line_request,omitempty"`
	// 账户类型。进项：AP，销项：AR
	AccountType string `json:"account_type,omitempty" xml:"account_type,omitempty"`
	// 含税金额
	AmountWithTax string `json:"amount_with_tax,omitempty" xml:"amount_with_tax,omitempty"`
	// 不含税金额
	AmountWithoutTax string `json:"amount_without_tax,omitempty" xml:"amount_without_tax,omitempty"`
	// 认证备注
	AuthRemark string `json:"auth_remark,omitempty" xml:"auth_remark,omitempty"`
	// 认证状态
	AuthStatus string `json:"auth_status,omitempty" xml:"auth_status,omitempty"`
	// 认证方式
	AuthStyle string `json:"auth_style,omitempty" xml:"auth_style,omitempty"`
	// 认证所属期
	AuthTaxPeriod string `json:"auth_tax_period,omitempty" xml:"auth_tax_period,omitempty"`
	// 抵扣用途
	AuthUse string `json:"auth_use,omitempty" xml:"auth_use,omitempty"`
	// 业务单号
	BizOrderNo string `json:"biz_order_no,omitempty" xml:"biz_order_no,omitempty"`
	// 购方地址电话
	BuyerAddressTel string `json:"buyer_address_tel,omitempty" xml:"buyer_address_tel,omitempty"`
	// 购方银行账号
	BuyerBankAccount string `json:"buyer_bank_account,omitempty" xml:"buyer_bank_account,omitempty"`
	// 购方开户行名称 账号
	BuyerBankInfo string `json:"buyer_bank_info,omitempty" xml:"buyer_bank_info,omitempty"`
	// 购方银行名称
	BuyerBankName string `json:"buyer_bank_name,omitempty" xml:"buyer_bank_name,omitempty"`
	// 购方名称
	BuyerName string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// 购方税号
	BuyerTaxNo string `json:"buyer_tax_no,omitempty" xml:"buyer_tax_no,omitempty"`
	// 校验码
	CheckCode string `json:"check_code,omitempty" xml:"check_code,omitempty"`
	// 密文
	CipherText string `json:"cipher_text,omitempty" xml:"cipher_text,omitempty"`
	// 二维密文
	CipherTextQRCode string `json:"cipher_text_q_r_code,omitempty" xml:"cipher_text_q_r_code,omitempty"`
	// 开票日期
	DateIssued string `json:"date_issued,omitempty" xml:"date_issued,omitempty"`
	// 有效税额
	EffectiveTaxAmount string `json:"effective_tax_amount,omitempty" xml:"effective_tax_amount,omitempty"`
	// 认证业务日期
	ElConfirmDate string `json:"el_confirm_date,omitempty" xml:"el_confirm_date,omitempty"`
	// 底账同步时间
	ElSyncTime string `json:"el_sync_time,omitempty" xml:"el_sync_time,omitempty"`
	// 扩展属性
	ExtendedAttrs string `json:"extended_attrs,omitempty" xml:"extended_attrs,omitempty"`
	// 发票代码
	InvoiceCode string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty"`
	// 红蓝票标记
	InvoiceColor string `json:"invoice_color,omitempty" xml:"invoice_color,omitempty"`
	// 发票类型
	InvoiceType string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 发票号码
	InvoiceNo string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
	// 国税发票来源
	InvoiceSource string `json:"invoice_source,omitempty" xml:"invoice_source,omitempty"`
	// 发票样式
	InvoiceStyleType string `json:"invoice_style_type,omitempty" xml:"invoice_style_type,omitempty"`
	// 开票人
	Issuer string `json:"issuer,omitempty" xml:"issuer,omitempty"`
	// 机器编码
	MachineCode string `json:"machine_code,omitempty" xml:"machine_code,omitempty"`
	// 不抵扣原因
	NoAuthReason string `json:"no_auth_reason,omitempty" xml:"no_auth_reason,omitempty"`
	// 原始发票代码
	OriginalInvoiceCode string `json:"original_invoice_code,omitempty" xml:"original_invoice_code,omitempty"`
	// 原始发票号码
	OriginalInvoiceNo string `json:"original_invoice_no,omitempty" xml:"original_invoice_no,omitempty"`
	// 付款人
	Payee string `json:"payee,omitempty" xml:"payee,omitempty"`
	// PDF文件地址
	PdfUrl string `json:"pdf_url,omitempty" xml:"pdf_url,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 被红冲含税金额
	ReverseAmountWithTax string `json:"reverse_amount_with_tax,omitempty" xml:"reverse_amount_with_tax,omitempty"`
	// 被红冲不含税金额
	ReverseAmountWithoutTax string `json:"reverse_amount_without_tax,omitempty" xml:"reverse_amount_without_tax,omitempty"`
	// 红冲状态
	ReverseFlag string `json:"reverse_flag,omitempty" xml:"reverse_flag,omitempty"`
	// 被红冲税额
	ReverseTaxAmount string `json:"reverse_tax_amount,omitempty" xml:"reverse_tax_amount,omitempty"`
	// 审核人
	Reviewer string `json:"reviewer,omitempty" xml:"reviewer,omitempty"`
	// 发是否有销货清单
	SaleListFileFlag string `json:"sale_list_file_flag,omitempty" xml:"sale_list_file_flag,omitempty"`
	// 销方地址
	SellerAddress string `json:"seller_address,omitempty" xml:"seller_address,omitempty"`
	// 销方地址电话
	SellerAddressTel string `json:"seller_address_tel,omitempty" xml:"seller_address_tel,omitempty"`
	// 销方银行账号
	SellerBankAccount string `json:"seller_bank_account,omitempty" xml:"seller_bank_account,omitempty"`
	// 销方开户行名称 账号
	SellerBankInfo string `json:"seller_bank_info,omitempty" xml:"seller_bank_info,omitempty"`
	// 销方银行名称
	SellerBankName string `json:"seller_bank_name,omitempty" xml:"seller_bank_name,omitempty"`
	// 销方名称
	SellerName string `json:"seller_name,omitempty" xml:"seller_name,omitempty"`
	// 销方税号
	SellerTaxNo string `json:"seller_tax_no,omitempty" xml:"seller_tax_no,omitempty"`
	// 销方电话
	SellerTel string `json:"seller_tel,omitempty" xml:"seller_tel,omitempty"`
	// 系统来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 发票特殊类型标识
	SpecialType string `json:"special_type,omitempty" xml:"special_type,omitempty"`
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 税额
	TaxAmount string `json:"tax_amount,omitempty" xml:"tax_amount,omitempty"`
	// 国税发票来源
	TaxInvoiceSource string `json:"tax_invoice_source,omitempty" xml:"tax_invoice_source,omitempty"`
	// 税率
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 发送时间错
	SendTime int64 `json:"send_time,omitempty" xml:"send_time,omitempty"`
}

var poolPytLedgerSyncRequest = sync.Pool{
	New: func() any {
		return new(PytLedgerSyncRequest)
	},
}

// GetPytLedgerSyncRequest() 从对象池中获取PytLedgerSyncRequest
func GetPytLedgerSyncRequest() *PytLedgerSyncRequest {
	return poolPytLedgerSyncRequest.Get().(*PytLedgerSyncRequest)
}

// ReleasePytLedgerSyncRequest 释放PytLedgerSyncRequest
func ReleasePytLedgerSyncRequest(v *PytLedgerSyncRequest) {
	v.Details = v.Details[:0]
	v.AccountType = ""
	v.AmountWithTax = ""
	v.AmountWithoutTax = ""
	v.AuthRemark = ""
	v.AuthStatus = ""
	v.AuthStyle = ""
	v.AuthTaxPeriod = ""
	v.AuthUse = ""
	v.BizOrderNo = ""
	v.BuyerAddressTel = ""
	v.BuyerBankAccount = ""
	v.BuyerBankInfo = ""
	v.BuyerBankName = ""
	v.BuyerName = ""
	v.BuyerTaxNo = ""
	v.CheckCode = ""
	v.CipherText = ""
	v.CipherTextQRCode = ""
	v.DateIssued = ""
	v.EffectiveTaxAmount = ""
	v.ElConfirmDate = ""
	v.ElSyncTime = ""
	v.ExtendedAttrs = ""
	v.InvoiceCode = ""
	v.InvoiceColor = ""
	v.InvoiceType = ""
	v.InvoiceNo = ""
	v.InvoiceSource = ""
	v.InvoiceStyleType = ""
	v.Issuer = ""
	v.MachineCode = ""
	v.NoAuthReason = ""
	v.OriginalInvoiceCode = ""
	v.OriginalInvoiceNo = ""
	v.Payee = ""
	v.PdfUrl = ""
	v.Remark = ""
	v.ReverseAmountWithTax = ""
	v.ReverseAmountWithoutTax = ""
	v.ReverseFlag = ""
	v.ReverseTaxAmount = ""
	v.Reviewer = ""
	v.SaleListFileFlag = ""
	v.SellerAddress = ""
	v.SellerAddressTel = ""
	v.SellerBankAccount = ""
	v.SellerBankInfo = ""
	v.SellerBankName = ""
	v.SellerName = ""
	v.SellerTaxNo = ""
	v.SellerTel = ""
	v.Source = ""
	v.SpecialType = ""
	v.Status = ""
	v.TaxAmount = ""
	v.TaxInvoiceSource = ""
	v.TaxRate = ""
	v.SendTime = 0
	poolPytLedgerSyncRequest.Put(v)
}
