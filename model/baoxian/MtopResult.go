package baoxian

// MtopResult 
type MtopResult struct {

    // model
    Model   string `json:"model,omitempty"`

    // errorCode
    ErrorCode   string `json:"error_code,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

    // isSuccess
    IsSuccess   bool `json:"is_success,omitempty"`

}