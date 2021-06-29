package maitix

// AlibabaDamaiMaitixDistributionDeliveryCalculateResult 
type AlibabaDamaiMaitixDistributionDeliveryCalculateResult struct {
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误码
    ErrorCode   int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 业务对象
    Model   *OpenApiPostFeeDto `json:"model,omitempty" xml:"model,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}