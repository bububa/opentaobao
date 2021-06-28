package promotion

// AlibabaLatourStrategyShowResult 
/* model for simplify = false
type AlibabaLatourStrategyShowResult struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 返回结果
    
    Data  *struct {
        StrategyShowResultDto  *StrategyShowResultDto `json:"strategy_show_result_dto,omitempty"`
    } `json:"data,omitempty"`
    

    // 错误描述
    
    Msg   string `json:"msg,omitempty"`
    

    // 错误码
    
    Code   string `json:"code,omitempty"`
    

}
*/

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
