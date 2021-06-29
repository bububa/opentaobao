package maitix

// OpenResult 
type OpenResult struct {

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 错误描述
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

    // 返回结果
    
    Model   *DisEncrypt4CmbResult `json:"model,omitempty" xml:"model,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 结果
    
    ModelList   []OpenExchangePointDto `json:"model_list,omitempty" xml:"model_list,omitempty"`
    

    // 参数extMap
    
    ExtMap   string `json:"ext_map,omitempty" xml:"ext_map,omitempty"`
    

    // 电子票对象
    
    ModelArrList   []EticketDto `json:"model_arr_list,omitempty" xml:"model_arr_list,omitempty"`
    

}
