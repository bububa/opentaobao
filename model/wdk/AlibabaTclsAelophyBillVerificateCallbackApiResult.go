package wdk

// AlibabaTclsAelophyBillVerificateCallbackApiResult 
type AlibabaTclsAelophyBillVerificateCallbackApiResult struct {

    // 回调是否处理成功
    Success   bool `json:"success,omitempty"`

    // 错误说明
    ErrMsg   string `json:"err_msg,omitempty"`

    // 错误编码
    ErrCode   string `json:"err_code,omitempty"`

}
