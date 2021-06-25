package aliqin

// AlibabaIsvDigitalsmsCreatetemplateResult 
type AlibabaIsvDigitalsmsCreatetemplateResult struct {

    // 模板code
    Model   string `json:"model,omitempty"`

    // 返回信息描述
    Msg   string `json:"msg,omitempty"`

    // true表示成功，false表示失败
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrCode   string `json:"err_code,omitempty"`

}