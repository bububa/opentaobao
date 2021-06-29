package lstvending

// AlibabaLstVendingTradeflowQueryResultDto 
type AlibabaLstVendingTradeflowQueryResultDto struct {

    // 错误信息
    
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    

    // 交易流水记录
    
    ModuleList   []VendingTradeFlowDto `json:"module_list,omitempty" xml:"module_list,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 是否异常
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
