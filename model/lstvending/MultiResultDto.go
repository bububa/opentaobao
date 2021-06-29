package lstvending

// MultiResultDto 
type MultiResultDto struct {

    // 错误信息
    
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    

    // 执行成功结果集
    
    ModuleList   []VendingCargoSpaceDto `json:"module_list,omitempty" xml:"module_list,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 执行失败结果集
    
    ErrorList   []AlibabaLstVendingCargospaceSaveResultDto `json:"error_list,omitempty" xml:"error_list,omitempty"`
    

    // 执行是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
