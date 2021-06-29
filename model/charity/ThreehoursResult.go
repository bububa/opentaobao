package charity

// ThreehoursResult 
type ThreehoursResult struct {

    // 消息
    
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
    

    // 错误码
    
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    

    // 结果
    
    Data   *UserActionSyncResult `json:"data,omitempty" xml:"data,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
