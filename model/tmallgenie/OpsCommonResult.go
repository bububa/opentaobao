package tmallgenie

// OpsCommonResult 
type OpsCommonResult struct {

    // 错误码，0为成功，其余为失败
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
    

    // 错误信息
    
    RetMsg   string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
    

}
