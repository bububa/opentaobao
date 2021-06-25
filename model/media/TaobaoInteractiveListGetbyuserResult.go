package media

// TaobaoInteractiveListGetbyuserResult 
type TaobaoInteractiveListGetbyuserResult struct {

    // 错误
    ResultCode   *ResultCode `json:"result_code,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // model
    Model   *PageQueryResult `json:"model,omitempty"`

}
