package traveltrade

// AlitripTravelTradeServiceinfoWriteResultSet 
type AlitripTravelTradeServiceinfoWriteResultSet struct {
    // 错误码信息
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 错误详细信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 是否操作成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 订单标注服务信息是否成功
    Error   bool `json:"error,omitempty" xml:"error,omitempty"`
}
