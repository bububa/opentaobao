package idle

// AlibabaIdleConsignmentOrderGetResult 
type AlibabaIdleConsignmentOrderGetResult struct {
    // 帮卖订单详情
    Module   *ConsignmentOrderTo `json:"module,omitempty" xml:"module,omitempty"`
    // 错误编码
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
