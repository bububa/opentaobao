package mos

// AlibabaMosOnsiteTradeQueryrefundResultDo 
/* model for simplify = false
type AlibabaMosOnsiteTradeQueryrefundResultDo struct {

    // data
    
    Data  *struct {
        OnsiteRefundResponse  *OnsiteRefundResponse `json:"onsite_refund_response,omitempty"`
    } `json:"data,omitempty"`
    

    // errCode
    
    ErrCode   int64 `json:"err_code,omitempty"`
    

    // errMsg
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

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
