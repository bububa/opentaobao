package aesolution

import (
	"sync"
)

// GlobalAeopTpLoanInfoDto 结构体
type GlobalAeopTpLoanInfoDto struct {
	// loan time
	LoanTime string `json:"loan_time,omitempty" xml:"loan_time,omitempty"`
	// loan amount
	LoanAmount *GlobalMoneyStr `json:"loan_amount,omitempty" xml:"loan_amount,omitempty"`
}

var poolGlobalAeopTpLoanInfoDto = sync.Pool{
	New: func() any {
		return new(GlobalAeopTpLoanInfoDto)
	},
}

// GetGlobalAeopTpLoanInfoDto() 从对象池中获取GlobalAeopTpLoanInfoDto
func GetGlobalAeopTpLoanInfoDto() *GlobalAeopTpLoanInfoDto {
	return poolGlobalAeopTpLoanInfoDto.Get().(*GlobalAeopTpLoanInfoDto)
}

// ReleaseGlobalAeopTpLoanInfoDto 释放GlobalAeopTpLoanInfoDto
func ReleaseGlobalAeopTpLoanInfoDto(v *GlobalAeopTpLoanInfoDto) {
	v.LoanTime = ""
	v.LoanAmount = nil
	poolGlobalAeopTpLoanInfoDto.Put(v)
}
