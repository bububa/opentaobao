package alihouse

import (
	"sync"
)

// BusinessActivityBankCardDto 结构体
type BusinessActivityBankCardDto struct {
	// 子户
	SubBankCardNo string `json:"sub_bank_card_no,omitempty" xml:"sub_bank_card_no,omitempty"`
	// 银联号
	BankContactLine string `json:"bank_contact_line,omitempty" xml:"bank_contact_line,omitempty"`
	// 结算账户
	BankCardNo string `json:"bank_card_no,omitempty" xml:"bank_card_no,omitempty"`
	// 别名
	BankAccountAlias string `json:"bank_account_alias,omitempty" xml:"bank_account_alias,omitempty"`
	// 账户名称
	BankAccountName string `json:"bank_account_name,omitempty" xml:"bank_account_name,omitempty"`
	// 账户类型
	AccountType int64 `json:"account_type,omitempty" xml:"account_type,omitempty"`
	// ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolBusinessActivityBankCardDto = sync.Pool{
	New: func() any {
		return new(BusinessActivityBankCardDto)
	},
}

// GetBusinessActivityBankCardDto() 从对象池中获取BusinessActivityBankCardDto
func GetBusinessActivityBankCardDto() *BusinessActivityBankCardDto {
	return poolBusinessActivityBankCardDto.Get().(*BusinessActivityBankCardDto)
}

// ReleaseBusinessActivityBankCardDto 释放BusinessActivityBankCardDto
func ReleaseBusinessActivityBankCardDto(v *BusinessActivityBankCardDto) {
	v.SubBankCardNo = ""
	v.BankContactLine = ""
	v.BankCardNo = ""
	v.BankAccountAlias = ""
	v.BankAccountName = ""
	v.AccountType = 0
	v.Id = 0
	poolBusinessActivityBankCardDto.Put(v)
}
