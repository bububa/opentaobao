package idle

// AlibabaIdleAppraiseOrderQueryResult 
type AlibabaIdleAppraiseOrderQueryResult struct {
    // errCode
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // errMsg
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // 订单信息
    Module   *AppraiseOrderInfoDTO `json:"module,omitempty" xml:"module,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
