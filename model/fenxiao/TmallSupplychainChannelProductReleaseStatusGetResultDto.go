package fenxiao

// TmallSupplychainChannelProductReleaseStatusGetResultDto 
type TmallSupplychainChannelProductReleaseStatusGetResultDto struct {
    // 执行结果
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 查询结果
    Module   *TopProductStatusResult `json:"module,omitempty" xml:"module,omitempty"`
    // 错误码
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
}
