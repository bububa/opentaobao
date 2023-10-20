package charity

import (
	"sync"
)

// CsrInvoiceAntChainSyncDto 结构体
type CsrInvoiceAntChainSyncDto struct {
	// 项目金额集合
	ProjectList []CsrInvoiceApplyProjectDto `json:"project_list,omitempty" xml:"project_list>csr_invoice_apply_project_dto,omitempty"`
	// 账单集合
	BillList []CsrInvoiceBillDto `json:"bill_list,omitempty" xml:"bill_list>csr_invoice_bill_dto,omitempty"`
	// 发票开票文件
	FileList []CsrInvoiceFileDto `json:"file_list,omitempty" xml:"file_list>csr_invoice_file_dto,omitempty"`
	// list
	BillNoList []CsrInvoiceBillNoDto `json:"bill_no_list,omitempty" xml:"bill_no_list>csr_invoice_bill_no_dto,omitempty"`
	// 社会统一信用代码
	UnifiedCreditCode string `json:"unified_credit_code,omitempty" xml:"unified_credit_code,omitempty"`
	// 机构id
	InvoiceOrgId string `json:"invoice_org_id,omitempty" xml:"invoice_org_id,omitempty"`
	// 联系人电话
	MerchantTel string `json:"merchant_tel,omitempty" xml:"merchant_tel,omitempty"`
	// 开户行账号
	AccountNumber string `json:"account_number,omitempty" xml:"account_number,omitempty"`
	// 联系人
	MerchantContact string `json:"merchant_contact,omitempty" xml:"merchant_contact,omitempty"`
	// 开户行
	AccountBank string `json:"account_bank,omitempty" xml:"account_bank,omitempty"`
	// 商家昵称
	MerchantNickName string `json:"merchant_nick_name,omitempty" xml:"merchant_nick_name,omitempty"`
	// 联系人地址
	MerchantAddress string `json:"merchant_address,omitempty" xml:"merchant_address,omitempty"`
	// 发票id
	InvoiceId string `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// 发票抬头
	InvoiceTitle string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// 备注-拒绝原因
	InvoiceRemark string `json:"invoice_remark,omitempty" xml:"invoice_remark,omitempty"`
	// 发票状态
	InvoiceState int64 `json:"invoice_state,omitempty" xml:"invoice_state,omitempty"`
	// 发票金额，分
	InvoiceAmount int64 `json:"invoice_amount,omitempty" xml:"invoice_amount,omitempty"`
	// 申请开票时间，ms
	ApplicationTime int64 `json:"application_time,omitempty" xml:"application_time,omitempty"`
	// 商家id
	MerchantId int64 `json:"merchant_id,omitempty" xml:"merchant_id,omitempty"`
	// 店铺id
	ShopId int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
}

var poolCsrInvoiceAntChainSyncDto = sync.Pool{
	New: func() any {
		return new(CsrInvoiceAntChainSyncDto)
	},
}

// GetCsrInvoiceAntChainSyncDto() 从对象池中获取CsrInvoiceAntChainSyncDto
func GetCsrInvoiceAntChainSyncDto() *CsrInvoiceAntChainSyncDto {
	return poolCsrInvoiceAntChainSyncDto.Get().(*CsrInvoiceAntChainSyncDto)
}

// ReleaseCsrInvoiceAntChainSyncDto 释放CsrInvoiceAntChainSyncDto
func ReleaseCsrInvoiceAntChainSyncDto(v *CsrInvoiceAntChainSyncDto) {
	v.ProjectList = v.ProjectList[:0]
	v.BillList = v.BillList[:0]
	v.FileList = v.FileList[:0]
	v.BillNoList = v.BillNoList[:0]
	v.UnifiedCreditCode = ""
	v.InvoiceOrgId = ""
	v.MerchantTel = ""
	v.AccountNumber = ""
	v.MerchantContact = ""
	v.AccountBank = ""
	v.MerchantNickName = ""
	v.MerchantAddress = ""
	v.InvoiceId = ""
	v.InvoiceTitle = ""
	v.InvoiceRemark = ""
	v.InvoiceState = 0
	v.InvoiceAmount = 0
	v.ApplicationTime = 0
	v.MerchantId = 0
	v.ShopId = 0
	poolCsrInvoiceAntChainSyncDto.Put(v)
}
