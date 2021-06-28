package wdk

// AlibabaWdkSkuUpdateApiResults 
type AlibabaWdkSkuUpdateApiResults struct {

    // 错误码
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

    // 各条记录结果
    
    Models   []AlibabaWdkSkuUpdateApiResult `json:"models,omitempty" xml:"models,omitempty"`
    

    // 接口调用成功标志
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
