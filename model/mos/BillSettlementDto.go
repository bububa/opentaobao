package mos

import (
	"sync"
)

// BillSettlementDto 结构体
type BillSettlementDto struct {
	// 发票列表
	InvoiceDTOList []SettleInvoiceDto `json:"invoice_d_t_o_list,omitempty" xml:"invoice_d_t_o_list>settle_invoice_dto,omitempty"`
	// 行号
	LineNo string `json:"line_no,omitempty" xml:"line_no,omitempty"`
	// 结算单金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 结算行税率，普票的税率必须是0，专票的税率不能为0
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 业务细类代码
	BizSubModuleCode string `json:"biz_sub_module_code,omitempty" xml:"biz_sub_module_code,omitempty"`
	// 发票类型 SPECIAL(&#34;专票&#34;)，ORDINARY(&#34;普票&#34;)
	InvoiceType string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 供应商id（可空）
	SupplierNo string `json:"supplier_no,omitempty" xml:"supplier_no,omitempty"`
	// 供应商名称（可空
	SupplierName string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
	// 收款方开户省
	BankProvince string `json:"bank_province,omitempty" xml:"bank_province,omitempty"`
	// 收款方开户市
	BankCity string `json:"bank_city,omitempty" xml:"bank_city,omitempty"`
	// 收款方银行账号
	AccountNo string `json:"account_no,omitempty" xml:"account_no,omitempty"`
	// 收款方账号名称
	AccountName string `json:"account_name,omitempty" xml:"account_name,omitempty"`
	// 银行号
	BankCode string `json:"bank_code,omitempty" xml:"bank_code,omitempty"`
	// 收款方开户行
	BankName string `json:"bank_name,omitempty" xml:"bank_name,omitempty"`
	// 收款方开户支行
	BankBranchName string `json:"bank_branch_name,omitempty" xml:"bank_branch_name,omitempty"`
	// 收款方开户支行code
	BankBranchCode string `json:"bank_branch_code,omitempty" xml:"bank_branch_code,omitempty"`
	// 账号类型，COMPANY企业/PERSON个人
	AccountTypes string `json:"account_types,omitempty" xml:"account_types,omitempty"`
	// 联行号
	CnapsCode string `json:"cnaps_code,omitempty" xml:"cnaps_code,omitempty"`
	// 联系人
	Contactor string `json:"contactor,omitempty" xml:"contactor,omitempty"`
	// 联系方式：国际化区号 + 座机/手机
	Telephone string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 付款说明。该字段会通过银行传给供应商，过长时会自动截取
	Comments string `json:"comments,omitempty" xml:"comments,omitempty"`
	// 扩展
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// 受益部门coa
	DepartmentCoa string `json:"department_coa,omitempty" xml:"department_coa,omitempty"`
	// 区域科目段
	CityCoa string `json:"city_coa,omitempty" xml:"city_coa,omitempty"`
}

var poolBillSettlementDto = sync.Pool{
	New: func() any {
		return new(BillSettlementDto)
	},
}

// GetBillSettlementDto() 从对象池中获取BillSettlementDto
func GetBillSettlementDto() *BillSettlementDto {
	return poolBillSettlementDto.Get().(*BillSettlementDto)
}

// ReleaseBillSettlementDto 释放BillSettlementDto
func ReleaseBillSettlementDto(v *BillSettlementDto) {
	v.InvoiceDTOList = v.InvoiceDTOList[:0]
	v.LineNo = ""
	v.Amount = ""
	v.TaxRate = ""
	v.BizSubModuleCode = ""
	v.InvoiceType = ""
	v.SupplierNo = ""
	v.SupplierName = ""
	v.BankProvince = ""
	v.BankCity = ""
	v.AccountNo = ""
	v.AccountName = ""
	v.BankCode = ""
	v.BankName = ""
	v.BankBranchName = ""
	v.BankBranchCode = ""
	v.AccountTypes = ""
	v.CnapsCode = ""
	v.Contactor = ""
	v.Telephone = ""
	v.Comments = ""
	v.ExtendParams = ""
	v.DepartmentCoa = ""
	v.CityCoa = ""
	poolBillSettlementDto.Put(v)
}
