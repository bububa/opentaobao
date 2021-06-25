package mos

// AlibabaMosOnsiteTradeQueryrefundResultDo 
type AlibabaMosOnsiteTradeQueryrefundResultDo struct {

    // data
    Data   *OnsiteRefundResponse `json:"data,omitempty"`

    // errCode
    ErrCode   int64 `json:"err_code,omitempty"`

    // errMsg
    ErrMsg   string `json:"err_msg,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
