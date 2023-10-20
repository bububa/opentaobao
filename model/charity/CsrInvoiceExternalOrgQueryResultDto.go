package charity

import (
	"sync"
)

// CsrInvoiceExternalOrgQueryResultDto 结构体
type CsrInvoiceExternalOrgQueryResultDto struct {
	// 项目金额明细
	ProjectList []CsrInvoiceApplyProjectDto `json:"project_list,omitempty" xml:"project_list>csr_invoice_apply_project_dto,omitempty"`
	// 商家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 商家联系人电话
	ContactMobile string `json:"contact_mobile,omitempty" xml:"contact_mobile,omitempty"`
	// 统一社会信用代码
	UnifiedCreditCode string `json:"unified_credit_code,omitempty" xml:"unified_credit_code,omitempty"`
	// 商家联系人
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 受理操作人
	AcceptOperator string `json:"accept_operator,omitempty" xml:"accept_operator,omitempty"`
	// 店铺名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 机构唯一标识
	OrgPromoteId string `json:"org_promote_id,omitempty" xml:"org_promote_id,omitempty"`
	// 开户行名称
	AccountBank string `json:"account_bank,omitempty" xml:"account_bank,omitempty"`
	// 开户行账号
	AccountNo string `json:"account_no,omitempty" xml:"account_no,omitempty"`
	// 纸质:paper/电子electronic
	InvoiceType string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 商家联系人地址
	ContactAddress string `json:"contact_address,omitempty" xml:"contact_address,omitempty"`
	// 发票id
	InvoiceId string `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// 发票抬头
	InvoiceTitle string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// 账单信息
	BillInfo *CsrInvoiceBillDto `json:"bill_info,omitempty" xml:"bill_info,omitempty"`
	// 发票状态，1:已受理
	InvoiceState int64 `json:"invoice_state,omitempty" xml:"invoice_state,omitempty"`
	// 票据金额，单位分
	InvoiceAmount int64 `json:"invoice_amount,omitempty" xml:"invoice_amount,omitempty"`
	// 店铺id
	ShopId int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 申请时间-ms时间戳
	ApplyTime int64 `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
}

var poolCsrInvoiceExternalOrgQueryResultDto = sync.Pool{
	New: func() any {
		return new(CsrInvoiceExternalOrgQueryResultDto)
	},
}

// GetCsrInvoiceExternalOrgQueryResultDto() 从对象池中获取CsrInvoiceExternalOrgQueryResultDto
func GetCsrInvoiceExternalOrgQueryResultDto() *CsrInvoiceExternalOrgQueryResultDto {
	return poolCsrInvoiceExternalOrgQueryResultDto.Get().(*CsrInvoiceExternalOrgQueryResultDto)
}

// ReleaseCsrInvoiceExternalOrgQueryResultDto 释放CsrInvoiceExternalOrgQueryResultDto
func ReleaseCsrInvoiceExternalOrgQueryResultDto(v *CsrInvoiceExternalOrgQueryResultDto) {
	v.ProjectList = v.ProjectList[:0]
	v.SellerNick = ""
	v.ContactMobile = ""
	v.UnifiedCreditCode = ""
	v.ContactName = ""
	v.AcceptOperator = ""
	v.ShopName = ""
	v.OrgPromoteId = ""
	v.AccountBank = ""
	v.AccountNo = ""
	v.InvoiceType = ""
	v.ContactAddress = ""
	v.InvoiceId = ""
	v.InvoiceTitle = ""
	v.BillInfo = nil
	v.InvoiceState = 0
	v.InvoiceAmount = 0
	v.ShopId = 0
	v.ApplyTime = 0
	poolCsrInvoiceExternalOrgQueryResultDto.Put(v)
}
