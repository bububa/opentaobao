package cainiaohandover

// GlspResponse 
type GlspResponse struct {

    // 查询是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    

    // 错误信息
    
    ErrorInfo   *ErrorInfo `json:"error_info,omitempty" xml:"error_info,omitempty"`
    

    // 请求结果
    
    Result   *SolutionServiceResQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
    

}
