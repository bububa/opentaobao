package einvoice

// TaxOptimizationSalaryPayResultQueryResultDto 
type TaxOptimizationSalaryPayResultQueryResultDto struct {
    // 结果列表
    SalaryDetailList   []SalaryDetailDto `json:"salary_detail_list,omitempty" xml:"salary_detail_list>salary_detail_dto,omitempty"`
}
