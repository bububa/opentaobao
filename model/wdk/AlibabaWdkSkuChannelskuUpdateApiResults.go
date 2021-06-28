package wdk

// AlibabaWdkSkuChannelskuUpdateApiResults 
type AlibabaWdkSkuChannelskuUpdateApiResults struct {

    // 单个商品返回结果集合
    
    Models   []AlibabaWdkSkuChannelskuUpdateApiResult `json:"models,omitempty" xml:"models,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // errCode
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // errMsg
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

}
