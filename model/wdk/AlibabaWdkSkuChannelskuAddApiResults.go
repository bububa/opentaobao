package wdk

// AlibabaWdkSkuChannelskuAddApiResults 
/* model for simplify = false
type AlibabaWdkSkuChannelskuAddApiResults struct {

    // 错误编码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 成功失败
    
    Success   bool `json:"success,omitempty"`
    

    // 返会结果集合
    
    Models  struct {
        AlibabaWdkSkuChannelskuAddApiResult  []AlibabaWdkSkuChannelskuAddApiResult `json:"alibaba_wdk_sku_channelsku_add_api_result,omitempty"`
    } `json:"models,omitempty"`
    

}
*/

// AlibabaWdkSkuChannelskuAddApiResults 
type AlibabaWdkSkuChannelskuAddApiResults struct {

    // 错误编码
    ErrCode   string `json:"err_code,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 成功失败
    Success   bool `json:"success,omitempty"`

    // 返会结果集合
    Models   []AlibabaWdkSkuChannelskuAddApiResult `json:"models,omitempty"`

}
