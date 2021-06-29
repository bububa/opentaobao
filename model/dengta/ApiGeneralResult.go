package dengta

// ApiGeneralResult 
type ApiGeneralResult struct {

    // 错误码
    
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    

    // 是否成功
    
    Value   bool `json:"value,omitempty" xml:"value,omitempty"`
    

    // 链路id
    
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    

    // 错误信息
    
    Message   *ReturnMessage `json:"message,omitempty" xml:"message,omitempty"`
    

}
