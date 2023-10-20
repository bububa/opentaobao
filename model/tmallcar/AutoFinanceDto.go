package tmallcar

import (
	"sync"
)

// AutoFinanceDto 结构体
type AutoFinanceDto struct {
	// 金融方案名称
	LoanPlanName string `json:"loan_plan_name,omitempty" xml:"loan_plan_name,omitempty"`
	// 金融方案英文名称
	LoanPlanEnName string `json:"loan_plan_en_name,omitempty" xml:"loan_plan_en_name,omitempty"`
	// 合同利率
	ContractInterestRate string `json:"contract_interest_rate,omitempty" xml:"contract_interest_rate,omitempty"`
	// 产品的基本利率
	LoanInterestRate string `json:"loan_interest_rate,omitempty" xml:"loan_interest_rate,omitempty"`
	// 首付比例
	DownPaymentRatio string `json:"down_payment_ratio,omitempty" xml:"down_payment_ratio,omitempty"`
	// 期数
	Periods int64 `json:"periods,omitempty" xml:"periods,omitempty"`
	// 首付金额（单位：分）
	DownPaymentAmount int64 `json:"down_payment_amount,omitempty" xml:"down_payment_amount,omitempty"`
	// 申请贷款金额（单位：分）
	LoanAmount int64 `json:"loan_amount,omitempty" xml:"loan_amount,omitempty"`
	// 月供（单位：分）
	MonthPayment int64 `json:"month_payment,omitempty" xml:"month_payment,omitempty"`
}

var poolAutoFinanceDto = sync.Pool{
	New: func() any {
		return new(AutoFinanceDto)
	},
}

// GetAutoFinanceDto() 从对象池中获取AutoFinanceDto
func GetAutoFinanceDto() *AutoFinanceDto {
	return poolAutoFinanceDto.Get().(*AutoFinanceDto)
}

// ReleaseAutoFinanceDto 释放AutoFinanceDto
func ReleaseAutoFinanceDto(v *AutoFinanceDto) {
	v.LoanPlanName = ""
	v.LoanPlanEnName = ""
	v.ContractInterestRate = ""
	v.LoanInterestRate = ""
	v.DownPaymentRatio = ""
	v.Periods = 0
	v.DownPaymentAmount = 0
	v.LoanAmount = 0
	v.MonthPayment = 0
	poolAutoFinanceDto.Put(v)
}
