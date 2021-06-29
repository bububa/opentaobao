package tmallchannel

// TmallChannelTradeApplyorderRefuseResultDto 
type TmallChannelTradeApplyorderRefuseResultDto struct {
    // 错误信息
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // 错误码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 是否执行成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
