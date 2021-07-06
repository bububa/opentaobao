package aesolution

// GlobalAeopTpLoanInfoDto 结构体
type GlobalAeopTpLoanInfoDto struct {
	// loan time
	LoanTime string `json:"loan_time,omitempty" xml:"loan_time,omitempty"`
	// loan amount
	LoanAmount *GlobalMoneyStr `json:"loan_amount,omitempty" xml:"loan_amount,omitempty"`
}
