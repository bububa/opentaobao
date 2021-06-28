package wdk

// AlibabaWdkSkuChannelskuQueryApiResult 
type AlibabaWdkSkuChannelskuQueryApiResult struct {

    // 业务调用是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 异常状态码
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // 业务数据模型
    
    Model   *ChannelSkuDo `json:"model,omitempty" xml:"model,omitempty"`
    

    // 异常信息
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

}
