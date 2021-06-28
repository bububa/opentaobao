package wdk

// FulfillSingleResult 
/* model for simplify = false
type FulfillSingleResult struct {

    // 处理结果
    
    Result   bool `json:"result,omitempty"`
    

    // 是否业务异常
    
    BizException   bool `json:"biz_exception,omitempty"`
    

    // 异常信息
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // 异常描述
    
    ErrorDesc   string `json:"error_desc,omitempty"`
    

    // 异常code
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// FulfillSingleResult 
type FulfillSingleResult struct {

    // 处理结果
    Result   bool `json:"result,omitempty"`

    // 是否业务异常
    BizException   bool `json:"biz_exception,omitempty"`

    // 异常信息
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 异常描述
    ErrorDesc   string `json:"error_desc,omitempty"`

    // 异常code
    ErrorCode   string `json:"error_code,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}
