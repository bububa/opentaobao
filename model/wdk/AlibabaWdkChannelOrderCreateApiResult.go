package wdk

// AlibabaWdkChannelOrderCreateApiResult 
type AlibabaWdkChannelOrderCreateApiResult struct {

    // 错误编码
    ErrCode   string `json:"err_code,omitempty"`

    // 返回内容
    Model   *OrderOperateResult `json:"model,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}