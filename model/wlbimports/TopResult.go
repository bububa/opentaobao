package wlbimports

// TopResult 
type TopResult struct {

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 面单信息
    Result   *WayBillDTO `json:"result,omitempty"`

    // 子错误信息
    SubErrorCode   string `json:"sub_error_code,omitempty"`

    // 子错误码
    SubErrorMsg   string `json:"sub_error_msg,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 结果数
    TotalResults   int64 `json:"total_results,omitempty"`

}
