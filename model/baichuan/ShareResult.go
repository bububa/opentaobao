package baichuan

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
