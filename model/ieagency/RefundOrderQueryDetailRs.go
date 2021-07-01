package ieagency

// RefundOrderQueryDetailRS 
type RefundOrderQueryDetailRS struct {
    // 错误码
    ApiErrorCode   int64 `json:"api_error_code,omitempty" xml:"api_error_code,omitempty"`
    // 错误描述
    ApiErrorMsg   string `json:"api_error_msg,omitempty" xml:"api_error_msg,omitempty"`
    // 错误trace
    ErrTrace   string `json:"err_trace,omitempty" xml:"err_trace,omitempty"`
    // 失败
    Failure   bool `json:"failure,omitempty" xml:"failure,omitempty"`
    // 机器名称
    HostName   string `json:"host_name,omitempty" xml:"host_name,omitempty"`
    // 申请单详情
    RefundOrderVo   *RefundOrderVo `json:"refund_order_vo,omitempty" xml:"refund_order_vo,omitempty"`
    // 成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
