package antifraud

// ResultWrapper 
type ResultWrapper struct {

    // 返回风控结果信息
    
    Result   *CollinadataQueryResult `json:"result,omitempty" xml:"result,omitempty"`
    

    // 返回结果码 0为成功,其他为错误码
    
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    

    // 返回是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 描述
    
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
    

}
