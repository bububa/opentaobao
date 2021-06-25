package fenxiao

// TmallSupplychainChannelProductQuantityUpdateResultDto 
type TmallSupplychainChannelProductQuantityUpdateResultDto struct {

    // 执行结果
    Success   bool `json:"success,omitempty"`

    // 更新内容
    Module   *TopProductQuantityResult `json:"module,omitempty"`

    // 错误码
    ErrorMessage   string `json:"error_message,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

}
