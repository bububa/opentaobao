package wdk

// AlibabaWdkSkuQueryApiResults 
type AlibabaWdkSkuQueryApiResults struct {

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 结果集合
    Models   []AlibabaWdkSkuQueryApiResult `json:"models,omitempty"`

    // 错误编码
    ErrCode   string `json:"err_code,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

}