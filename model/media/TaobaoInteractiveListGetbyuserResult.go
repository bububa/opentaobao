package media

// TaobaoInteractiveListGetbyuserResult 
/* model for simplify = false
type TaobaoInteractiveListGetbyuserResult struct {

    // 错误
    
    ResultCode  *struct {
        ResultCode  *ResultCode `json:"result_code,omitempty"`
    } `json:"result_code,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // model
    
    Model  *struct {
        PageQueryResult  *PageQueryResult `json:"page_query_result,omitempty"`
    } `json:"model,omitempty"`
    

}
*/

// TaobaoInteractiveListGetbyuserResult 
type TaobaoInteractiveListGetbyuserResult struct {

    // 错误
    ResultCode   *ResultCode `json:"result_code,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // model
    Model   *PageQueryResult `json:"model,omitempty"`

}
