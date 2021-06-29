package wdk

// AlibabaWdkSkuCategoryQueryApiResult 
type AlibabaWdkSkuCategoryQueryApiResult struct {
    // 接口返回成功标志
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误码（只有有异常才有值）
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 错误信息（只有有异常才有值）
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // 返回业务参数模型，-1状态的为删除的类目
    Model   string `json:"model,omitempty" xml:"model,omitempty"`
}
