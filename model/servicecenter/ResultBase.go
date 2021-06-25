package servicecenter

// ResultBase 
type ResultBase struct {

    // value
    Value   *TpFundsRecoverResultDo `json:"value,omitempty"`

    // 查询接口是否OK
    Success   bool `json:"success,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

    // errorCode
    ErrorCode   int64 `json:"error_code,omitempty"`

}
