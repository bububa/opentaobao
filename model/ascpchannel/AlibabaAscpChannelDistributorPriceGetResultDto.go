package ascpchannel

// AlibabaAscpChannelDistributorPriceGetResultDTO 
type AlibabaAscpChannelDistributorPriceGetResultDTO struct {
    // 执行结果
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误码
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 价格数据
    Data   *TopDistributorPriceResult `json:"data,omitempty" xml:"data,omitempty"`
}
