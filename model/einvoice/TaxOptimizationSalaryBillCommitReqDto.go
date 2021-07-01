package einvoice

// TaxOptimizationSalaryBillCommitReqDto 
type TaxOptimizationSalaryBillCommitReqDto struct {
    // 账期
    AccountDate   string `json:"account_date,omitempty" xml:"account_date,omitempty"`
    // 是否关闭账期
    CloseAccountDate   bool `json:"close_account_date,omitempty" xml:"close_account_date,omitempty"`
    // 发薪isv对应的发薪额度
    ContractorAppliedDutiableAmount   string `json:"contractor_applied_dutiable_amount,omitempty" xml:"contractor_applied_dutiable_amount,omitempty"`
    // 发薪明细
    DetailList   []TaxOptimizationSalaryDetailInfoDto `json:"detail_list,omitempty" xml:"detail_list>tax_optimization_salary_detail_info_dto,omitempty"`
    // 业务方编码
    EmployerCode   string `json:"employer_code,omitempty" xml:"employer_code,omitempty"`
    // 是否开启账单
    StartAccountDate   bool `json:"start_account_date,omitempty" xml:"start_account_date,omitempty"`
    // 总账单明细数
    TotalDetailCount   int64 `json:"total_detail_count,omitempty" xml:"total_detail_count,omitempty"`
}
