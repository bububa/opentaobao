package mos

// AlibabaMosOnsiteTradeIsnewpayorderResultDo 
/* model for simplify = false
type AlibabaMosOnsiteTradeIsnewpayorderResultDo struct {

    // 是否为新支付订单。true：是，false:不是
    
    Data   bool `json:"data,omitempty"`
    

    // errCode
    
    ErrCode   int64 `json:"err_code,omitempty"`
    

    // errMsg
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlibabaMosOnsiteTradeIsnewpayorderResultDo 
type AlibabaMosOnsiteTradeIsnewpayorderResultDo struct {

    // 是否为新支付订单。true：是，false:不是
    Data   bool `json:"data,omitempty"`

    // errCode
    ErrCode   int64 `json:"err_code,omitempty"`

    // errMsg
    ErrMsg   string `json:"err_msg,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
