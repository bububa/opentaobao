package jym

// AlibabaJymRequirementOrdersOperationNotifyResultDto 
type AlibabaJymRequirementOrdersOperationNotifyResultDto struct {

    // 业务数据
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
    

    // 执行结果
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 状态码
    
    StateCode   string `json:"state_code,omitempty" xml:"state_code,omitempty"`
    

    // 错误信息
    
    ExtraErrMsg   string `json:"extra_err_msg,omitempty" xml:"extra_err_msg,omitempty"`
    

}
