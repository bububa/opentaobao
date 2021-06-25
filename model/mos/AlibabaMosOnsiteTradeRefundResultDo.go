package mos

// AlibabaMosOnsiteTradeRefundResultDo 
type AlibabaMosOnsiteTradeRefundResultDo struct {

    // data
    Data   *RefundResponse `json:"data,omitempty"`

    // 错误码
    ErrCode   int64 `json:"err_code,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
