package einvoice

// TaxOptimizationSalaryPaymentAccessResultDto 
type TaxOptimizationSalaryPaymentAccessResultDto struct {

    // 失败的个数
    
    FailCount   int64 `json:"fail_count,omitempty" xml:"fail_count,omitempty"`
    

    // 正在处理的个数
    
    ProcessingCount   int64 `json:"processing_count,omitempty" xml:"processing_count,omitempty"`
    

    // 发薪状态
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 成功的个数
    
    SuccessCount   int64 `json:"success_count,omitempty" xml:"success_count,omitempty"`
    

    // 总的发薪个数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    

}
