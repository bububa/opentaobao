package idle

// AlibabaIdleAgreementPayResult 
type AlibabaIdleAgreementPayResult struct {
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误code
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 服务出参
    Module   *Serializable `json:"module,omitempty" xml:"module,omitempty"`
    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}
