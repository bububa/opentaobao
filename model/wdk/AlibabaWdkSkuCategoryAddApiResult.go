package wdk

// AlibabaWdkSkuCategoryAddApiResult 
type AlibabaWdkSkuCategoryAddApiResult struct {

    // 接口返回成功标志
    Success   bool `json:"success,omitempty"`

    // 错误码（只有有异常才有值）
    ErrCode   string `json:"err_code,omitempty"`

    // 错误信息（只有有异常才有值）
    ErrMsg   string `json:"err_msg,omitempty"`

    // 调用成功时的返回类目code
    Model   string `json:"model,omitempty"`

}
