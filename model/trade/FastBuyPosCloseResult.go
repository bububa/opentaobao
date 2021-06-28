package trade

// FastBuyPosCloseResult 
type FastBuyPosCloseResult struct {

    // 关单结果状态: 1为成功,2为失败
    
    ResultResult   int64 `json:"result_result,omitempty" xml:"result_result,omitempty"`
    

    // 错误码
    
    ReturnCode   string `json:"return_code,omitempty" xml:"return_code,omitempty"`
    

    // 错误信息
    
    ReturnMsg   string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
    

    // 接口调用是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
