package promotion

// AlibabaLatourStrategyShowResult 
type AlibabaLatourStrategyShowResult struct {

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 返回结果
    Data   *StrategyShowResultDto `json:"data,omitempty"`

    // 错误描述
    Msg   string `json:"msg,omitempty"`

    // 错误码
    Code   string `json:"code,omitempty"`

}
