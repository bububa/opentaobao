package mos

import (
	"sync"
)

// SupplierBankInfoDto 结构体
type SupplierBankInfoDto struct {
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 国家名
	CountryName string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	// 地区名
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 省代码
	ProvinceCode string `json:"province_code,omitempty" xml:"province_code,omitempty"`
	// 城市代码
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 省名
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 门店号
	CoCode string `json:"co_code,omitempty" xml:"co_code,omitempty"`
	// 收款银行账号
	ReceiveAccount string `json:"receive_account,omitempty" xml:"receive_account,omitempty"`
	// 支行名
	BankBranchName string `json:"bank_branch_name,omitempty" xml:"bank_branch_name,omitempty"`
	// 联行号
	CnapsCode string `json:"cnaps_code,omitempty" xml:"cnaps_code,omitempty"`
	// 收款账户名
	ReceiveName string `json:"receive_name,omitempty" xml:"receive_name,omitempty"`
	// 门店名
	CoName string `json:"co_name,omitempty" xml:"co_name,omitempty"`
	// 支行号
	BankBranchCode string `json:"bank_branch_code,omitempty" xml:"bank_branch_code,omitempty"`
	// 供应商号
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 国家号
	CountryCode string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 银行名
	BankName string `json:"bank_name,omitempty" xml:"bank_name,omitempty"`
	// 银行号
	BankCode string `json:"bank_code,omitempty" xml:"bank_code,omitempty"`
	// 主键
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 审核状态
	ApprovalStatus int64 `json:"approval_status,omitempty" xml:"approval_status,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 业务乐行
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 关联账号id
	RelationAccountId int64 `json:"relation_account_id,omitempty" xml:"relation_account_id,omitempty"`
	// 类型, 个人, 企业
	AccountType int64 `json:"account_type,omitempty" xml:"account_type,omitempty"`
	// 操作类型
	ApprovalType int64 `json:"approval_type,omitempty" xml:"approval_type,omitempty"`
}

var poolSupplierBankInfoDto = sync.Pool{
	New: func() any {
		return new(SupplierBankInfoDto)
	},
}

// GetSupplierBankInfoDto() 从对象池中获取SupplierBankInfoDto
func GetSupplierBankInfoDto() *SupplierBankInfoDto {
	return poolSupplierBankInfoDto.Get().(*SupplierBankInfoDto)
}

// ReleaseSupplierBankInfoDto 释放SupplierBankInfoDto
func ReleaseSupplierBankInfoDto(v *SupplierBankInfoDto) {
	v.GmtModified = ""
	v.CountryName = ""
	v.CityName = ""
	v.ProvinceCode = ""
	v.CityCode = ""
	v.ProvinceName = ""
	v.CoCode = ""
	v.ReceiveAccount = ""
	v.BankBranchName = ""
	v.CnapsCode = ""
	v.ReceiveName = ""
	v.CoName = ""
	v.BankBranchCode = ""
	v.SupplierId = ""
	v.GmtCreate = ""
	v.CountryCode = ""
	v.BankName = ""
	v.BankCode = ""
	v.Id = 0
	v.ApprovalStatus = 0
	v.Status = 0
	v.BusinessType = 0
	v.RelationAccountId = 0
	v.AccountType = 0
	v.ApprovalType = 0
	poolSupplierBankInfoDto.Put(v)
}
