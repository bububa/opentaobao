package txcs

// CommonResult 
type CommonResult struct {

    // 错误码
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // 创建的互动实例
    
    Data   *StatementBillDto `json:"data,omitempty" xml:"data,omitempty"`
    

    // 服务标识
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 错误信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 服务标识
    
    BizException   bool `json:"biz_exception,omitempty" xml:"biz_exception,omitempty"`
    

}
