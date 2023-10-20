package alihouse

import (
	"sync"
)

// AccountDivisionRuleDto 结构体
type AccountDivisionRuleDto struct {
	// 比例
	Ratio string `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// 联行号
	AccountContactLine string `json:"account_contact_line,omitempty" xml:"account_contact_line,omitempty"`
	// 开户行名称
	AccountBankName string `json:"account_bank_name,omitempty" xml:"account_bank_name,omitempty"`
	// 银行账号
	AccountNo string `json:"account_no,omitempty" xml:"account_no,omitempty"`
	// 账户名
	AccountName string `json:"account_name,omitempty" xml:"account_name,omitempty"`
	// 外部分账ID
	OuterDivisionId string `json:"outer_division_id,omitempty" xml:"outer_division_id,omitempty"`
	// 账户类型
	AccountType string `json:"account_type,omitempty" xml:"account_type,omitempty"`
}

var poolAccountDivisionRuleDto = sync.Pool{
	New: func() any {
		return new(AccountDivisionRuleDto)
	},
}

// GetAccountDivisionRuleDto() 从对象池中获取AccountDivisionRuleDto
func GetAccountDivisionRuleDto() *AccountDivisionRuleDto {
	return poolAccountDivisionRuleDto.Get().(*AccountDivisionRuleDto)
}

// ReleaseAccountDivisionRuleDto 释放AccountDivisionRuleDto
func ReleaseAccountDivisionRuleDto(v *AccountDivisionRuleDto) {
	v.Ratio = ""
	v.AccountContactLine = ""
	v.AccountBankName = ""
	v.AccountNo = ""
	v.AccountName = ""
	v.OuterDivisionId = ""
	v.AccountType = ""
	poolAccountDivisionRuleDto.Put(v)
}
