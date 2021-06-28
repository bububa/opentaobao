package wdk

// AlibabaWdkSkuChannelskuQueryApiResults 
/* model for simplify = false
type AlibabaWdkSkuChannelskuQueryApiResults struct {

    // 错误编码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // skuCode不能为空
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 接口调用是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 业务数据模型
    
    Models  struct {
        AlibabaWdkSkuChannelskuQueryApiResult  []AlibabaWdkSkuChannelskuQueryApiResult `json:"alibaba_wdk_sku_channelsku_query_api_result,omitempty"`
    } `json:"models,omitempty"`
    

}
*/

// AlibabaWdkSkuChannelskuQueryApiResults 
type AlibabaWdkSkuChannelskuQueryApiResults struct {

    // 错误编码
    ErrCode   string `json:"err_code,omitempty"`

    // skuCode不能为空
    ErrMsg   string `json:"err_msg,omitempty"`

    // 接口调用是否成功
    Success   bool `json:"success,omitempty"`

    // 业务数据模型
    Models   []AlibabaWdkSkuChannelskuQueryApiResult `json:"models,omitempty"`

}
