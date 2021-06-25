package security

// JaqAccountRiskResult 
type JaqAccountRiskResult struct {

    // 风控结果的真正数据内容
    Data   *JaqAccountRiskData `json:"data,omitempty"`

    // 错误码
    ErrorCode   int64 `json:"error_code,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

}
