package wdk

// FulfillSingleResult 
type FulfillSingleResult struct {

    // 处理结果
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
    

    // 是否业务异常
    
    BizException   bool `json:"biz_exception,omitempty" xml:"biz_exception,omitempty"`
    

    // 异常信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

    // 异常描述
    
    ErrorDesc   string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
    

    // 异常code
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
