package scbp

import (
	"sync"
)

// AccountReportDto 结构体
type AccountReportDto struct {
	// 返回数据集合
	AccountEffectList []AccountEffectDto `json:"account_effect_list,omitempty" xml:"account_effect_list>account_effect_dto,omitempty"`
}

var poolAccountReportDto = sync.Pool{
	New: func() any {
		return new(AccountReportDto)
	},
}

// GetAccountReportDto() 从对象池中获取AccountReportDto
func GetAccountReportDto() *AccountReportDto {
	return poolAccountReportDto.Get().(*AccountReportDto)
}

// ReleaseAccountReportDto 释放AccountReportDto
func ReleaseAccountReportDto(v *AccountReportDto) {
	v.AccountEffectList = v.AccountEffectList[:0]
	poolAccountReportDto.Put(v)
}
