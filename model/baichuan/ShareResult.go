package baichuan

// ShareResult 
/* model for simplify = false
type ShareResult struct {

    // model
    
    Model  *struct {
        PasswordRuleResultDto  *PasswordRuleResultDto `json:"password_rule_result_dto,omitempty"`
    } `json:"model,omitempty"`
    

    // resultCode
    
    ResultCode  *struct {
        ResultCode  *ResultCode `json:"result_code,omitempty"`
    } `json:"result_code,omitempty"`
    

    // totalNumber
    
    TotalNumber   int64 `json:"total_number,omitempty"`
    

    // isSuccess
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

// ShareResult 
type ShareResult struct {

    // model
    Model   *PasswordRuleResultDto `json:"model,omitempty"`

    // resultCode
    ResultCode   *ResultCode `json:"result_code,omitempty"`

    // totalNumber
    TotalNumber   int64 `json:"total_number,omitempty"`

    // isSuccess
    IsSuccess   bool `json:"is_success,omitempty"`

}
