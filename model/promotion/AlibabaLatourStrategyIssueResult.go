package promotion

// AlibabaLatourStrategyIssueResult 
type AlibabaLatourStrategyIssueResult struct {

    // 错误码
    Code   string `json:"code,omitempty"`

    // 返回结果
    Data   *StrategyIssueResultDto `json:"data,omitempty"`

    // 错误描述
    Msg   string `json:"msg,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}
