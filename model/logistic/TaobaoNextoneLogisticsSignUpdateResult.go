package logistic

// TaobaoNextoneLogisticsSignUpdateResult 
type TaobaoNextoneLogisticsSignUpdateResult struct {

    // 返回数据
    ResultData   string `json:"result_data,omitempty"`

    // 错误信息
    ErrorInfo   string `json:"error_info,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 成功失败
    Success   bool `json:"success,omitempty"`

}
