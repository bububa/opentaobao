package alihouse

import (
	"sync"
)

// ExistingHomeBankCardDto 结构体
type ExistingHomeBankCardDto struct {
	// 银联号
	BankContactLine string `json:"bank_contact_line,omitempty" xml:"bank_contact_line,omitempty"`
	// 结算账户
	BankCardNo string `json:"bank_card_no,omitempty" xml:"bank_card_no,omitempty"`
	// 结算账户别名
	BankAccountAlias string `json:"bank_account_alias,omitempty" xml:"bank_account_alias,omitempty"`
	// 结算账户名称
	BankAccountName string `json:"bank_account_name,omitempty" xml:"bank_account_name,omitempty"`
	// 外部账户ID
	OuterBankId string `json:"outer_bank_id,omitempty" xml:"outer_bank_id,omitempty"`
	// 账户类型 01对私账户 02对公账户
	AccountType string `json:"account_type,omitempty" xml:"account_type,omitempty"`
	// 银行名称简称
	BankNameSimple string `json:"bank_name_simple,omitempty" xml:"bank_name_simple,omitempty"`
	// 银行名称
	BankName string `json:"bank_name,omitempty" xml:"bank_name,omitempty"`
	// 支行所在省份
	BankBranchProvince string `json:"bank_branch_province,omitempty" xml:"bank_branch_province,omitempty"`
	// 支行所在市区
	BankBranchCity string `json:"bank_branch_city,omitempty" xml:"bank_branch_city,omitempty"`
	// 支行名称
	BankBranchName string `json:"bank_branch_name,omitempty" xml:"bank_branch_name,omitempty"`
	// 是否有效 1-是 0-否
	IsValid int64 `json:"is_valid,omitempty" xml:"is_valid,omitempty"`
	// 是否是四方监管户 1-是 0-否
	IsFourAccount int64 `json:"is_four_account,omitempty" xml:"is_four_account,omitempty"`
	// 账户用途
	AccountUseType int64 `json:"account_use_type,omitempty" xml:"account_use_type,omitempty"`
	// 公司进件ID
	MerchantOpenId int64 `json:"merchant_open_id,omitempty" xml:"merchant_open_id,omitempty"`
	// 卡类型
	CardType int64 `json:"card_type,omitempty" xml:"card_type,omitempty"`
	// 是否开网商子户
	IsOpenNetBusiness int64 `json:"is_open_net_business,omitempty" xml:"is_open_net_business,omitempty"`
	// 场景 5-融合店
	SceneType int64 `json:"scene_type,omitempty" xml:"scene_type,omitempty"`
}

var poolExistingHomeBankCardDto = sync.Pool{
	New: func() any {
		return new(ExistingHomeBankCardDto)
	},
}

// GetExistingHomeBankCardDto() 从对象池中获取ExistingHomeBankCardDto
func GetExistingHomeBankCardDto() *ExistingHomeBankCardDto {
	return poolExistingHomeBankCardDto.Get().(*ExistingHomeBankCardDto)
}

// ReleaseExistingHomeBankCardDto 释放ExistingHomeBankCardDto
func ReleaseExistingHomeBankCardDto(v *ExistingHomeBankCardDto) {
	v.BankContactLine = ""
	v.BankCardNo = ""
	v.BankAccountAlias = ""
	v.BankAccountName = ""
	v.OuterBankId = ""
	v.AccountType = ""
	v.BankNameSimple = ""
	v.BankName = ""
	v.BankBranchProvince = ""
	v.BankBranchCity = ""
	v.BankBranchName = ""
	v.IsValid = 0
	v.IsFourAccount = 0
	v.AccountUseType = 0
	v.MerchantOpenId = 0
	v.CardType = 0
	v.IsOpenNetBusiness = 0
	v.SceneType = 0
	poolExistingHomeBankCardDto.Put(v)
}
