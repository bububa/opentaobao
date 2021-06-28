package promotion

// AlibabaLatourStrategyIssueResult 
/* model for simplify = false
type AlibabaLatourStrategyIssueResult struct {

    // 错误码
    
    Code   string `json:"code,omitempty"`
    

    // 返回结果
    
    Data  *struct {
        StrategyIssueResultDto  *StrategyIssueResultDto `json:"strategy_issue_result_dto,omitempty"`
    } `json:"data,omitempty"`
    

    // 错误描述
    
    Msg   string `json:"msg,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

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
