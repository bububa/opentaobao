package mos

import (
	"sync"
)

// ModifyBillDto 结构体
type ModifyBillDto struct {
	// 结算行号list
	SettleLineNos []string `json:"settle_line_nos,omitempty" xml:"settle_line_nos>string,omitempty"`
	// 收款方开户支行code
	BankBranchCode string `json:"bank_branch_code,omitempty" xml:"bank_branch_code,omitempty"`
	// 供应商名称（可空）
	SupplierName string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
	// 供应商id（可空
	SupplierNo string `json:"supplier_no,omitempty" xml:"supplier_no,omitempty"`
	// 收款方开户市
	BankCity string `json:"bank_city,omitempty" xml:"bank_city,omitempty"`
	// 收款方银行账号
	AccountNo string `json:"account_no,omitempty" xml:"account_no,omitempty"`
	// 收款方开户行
	BankName string `json:"bank_name,omitempty" xml:"bank_name,omitempty"`
	// 账号类型，COMPANY企业/PERSON个人
	AccountTypes string `json:"account_types,omitempty" xml:"account_types,omitempty"`
	// 收款方开户省
	BankProvince string `json:"bank_province,omitempty" xml:"bank_province,omitempty"`
	// //银行号
	BankCode string `json:"bank_code,omitempty" xml:"bank_code,omitempty"`
	// 收款方账号名称
	AccountName string `json:"account_name,omitempty" xml:"account_name,omitempty"`
	// 收款方开户支行
	BankBranchName string `json:"bank_branch_name,omitempty" xml:"bank_branch_name,omitempty"`
	// 联行号
	CnapsCode string `json:"cnaps_code,omitempty" xml:"cnaps_code,omitempty"`
	// //单据请求唯一编号，必填
	BillNo string `json:"bill_no,omitempty" xml:"bill_no,omitempty"`
	// 扩展参数
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
}

var poolModifyBillDto = sync.Pool{
	New: func() any {
		return new(ModifyBillDto)
	},
}

// GetModifyBillDto() 从对象池中获取ModifyBillDto
func GetModifyBillDto() *ModifyBillDto {
	return poolModifyBillDto.Get().(*ModifyBillDto)
}

// ReleaseModifyBillDto 释放ModifyBillDto
func ReleaseModifyBillDto(v *ModifyBillDto) {
	v.SettleLineNos = v.SettleLineNos[:0]
	v.BankBranchCode = ""
	v.SupplierName = ""
	v.SupplierNo = ""
	v.BankCity = ""
	v.AccountNo = ""
	v.BankName = ""
	v.AccountTypes = ""
	v.BankProvince = ""
	v.BankCode = ""
	v.AccountName = ""
	v.BankBranchName = ""
	v.CnapsCode = ""
	v.BillNo = ""
	v.ExtendParams = ""
	poolModifyBillDto.Put(v)
}
