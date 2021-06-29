package wdk

// AlibabaWdkChannelOrderStatusUpdateApiResult 
type AlibabaWdkChannelOrderStatusUpdateApiResult struct {
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // 返回内容
    Model   *OrderOperateResult `json:"model,omitempty" xml:"model,omitempty"`
    // 错误编码
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
}
