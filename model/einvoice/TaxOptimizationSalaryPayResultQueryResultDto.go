package einvoice

// TaxOptimizationSalaryPayResultQueryResultDTO 
type TaxOptimizationSalaryPayResultQueryResultDTO struct {
    // 结果列表
    SalaryDetailList   []SalaryDetailDTO `json:"salary_detail_list,omitempty" xml:"salary_detail_list>salary_detail_dto,omitempty"`
}
