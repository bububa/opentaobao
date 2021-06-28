package wdk

// AlibabaWdkSkuChannelskuUpdateApiResults 
/* model for simplify = false
type AlibabaWdkSkuChannelskuUpdateApiResults struct {

    // 单个商品返回结果集合
    
    Models  struct {
        AlibabaWdkSkuChannelskuUpdateApiResult  []AlibabaWdkSkuChannelskuUpdateApiResult `json:"alibaba_wdk_sku_channelsku_update_api_result,omitempty"`
    } `json:"models,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // errCode
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // errMsg
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

}
*/

// AlibabaWdkSkuChannelskuUpdateApiResults 
type AlibabaWdkSkuChannelskuUpdateApiResults struct {

    // 单个商品返回结果集合
    Models   []AlibabaWdkSkuChannelskuUpdateApiResult `json:"models,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // errCode
    ErrCode   string `json:"err_code,omitempty"`

    // errMsg
    ErrMsg   string `json:"err_msg,omitempty"`

}
