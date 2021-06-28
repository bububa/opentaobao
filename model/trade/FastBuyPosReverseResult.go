package trade

// FastBuyPosReverseResult 
/* model for simplify = false
type FastBuyPosReverseResult struct {

    // 返回的错误码
    
    ReturnCode   string `json:"return_code,omitempty"`
    

    // 返回的错误信息
    
    ReturnMsg   string `json:"return_msg,omitempty"`
    

    // 调用接口是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 退款状态: 1为成功,2为失败
    
    ResultStatus   int64 `json:"result_status,omitempty"`
    

}
*/

// FastBuyPosReverseResult 
type FastBuyPosReverseResult struct {

    // 返回的错误码
    ReturnCode   string `json:"return_code,omitempty"`

    // 返回的错误信息
    ReturnMsg   string `json:"return_msg,omitempty"`

    // 调用接口是否成功
    Success   bool `json:"success,omitempty"`

    // 退款状态: 1为成功,2为失败
    ResultStatus   int64 `json:"result_status,omitempty"`

}
