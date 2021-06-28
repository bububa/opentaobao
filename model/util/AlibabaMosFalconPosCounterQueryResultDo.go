package util

// AlibabaMosFalconPosCounterQueryResultDo 
type AlibabaMosFalconPosCounterQueryResultDo struct {

    // traceId
    
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    

    // total
    
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    

    // 是否成功
    
    Data   *PosInfoDto `json:"data,omitempty" xml:"data,omitempty"`
    

    // 调用是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 错误码
    
    ErrCode   int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // 额外
    
    Extra   string `json:"extra,omitempty" xml:"extra,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

    // 标题
    
    Titles   []string `json:"titles,omitempty" xml:"titles>string,omitempty"`
    

}
