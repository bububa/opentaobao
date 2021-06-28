package trade

// CainiaoRefundRefundactionsDisplayBizResult 
/* model for simplify = false
type CainiaoRefundRefundactionsDisplayBizResult struct {

    // 调用时错误码
    
    StatusCode   string `json:"status_code,omitempty"`
    

    // 买家操作时需要的展示信息
    
    Data   string `json:"data,omitempty"`
    

    // true表示成功，false表示失败
    
    Success   bool `json:"success,omitempty"`
    

    // 调用时错误描述
    
    StatusMessage   string `json:"status_message,omitempty"`
    

}
*/

// CainiaoRefundRefundactionsDisplayBizResult 
type CainiaoRefundRefundactionsDisplayBizResult struct {

    // 调用时错误码
    StatusCode   string `json:"status_code,omitempty"`

    // 买家操作时需要的展示信息
    Data   string `json:"data,omitempty"`

    // true表示成功，false表示失败
    Success   bool `json:"success,omitempty"`

    // 调用时错误描述
    StatusMessage   string `json:"status_message,omitempty"`

}
