package antifraud

// ResultWrapper 
/* model for simplify = false
type ResultWrapper struct {

    // 返回风控结果信息
    
    Result  *struct {
        CollinadataQueryResult  *CollinadataQueryResult `json:"collinadata_query_result,omitempty"`
    } `json:"result,omitempty"`
    

    // 返回结果码 0为成功,其他为错误码
    
    Code   int64 `json:"code,omitempty"`
    

    // 返回是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 描述
    
    Msg   string `json:"msg,omitempty"`
    

}
*/

// ResultWrapper 
type ResultWrapper struct {

    // 返回风控结果信息
    Result   *CollinadataQueryResult `json:"result,omitempty"`

    // 返回结果码 0为成功,其他为错误码
    Code   int64 `json:"code,omitempty"`

    // 返回是否成功
    Success   bool `json:"success,omitempty"`

    // 描述
    Msg   string `json:"msg,omitempty"`

}
