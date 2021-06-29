package aliexpresssumaitong

// ResponseDto 
type ResponseDto struct {

    // 数据
    
    Data   *HjTaxCalculateResultDto `json:"data,omitempty" xml:"data,omitempty"`
    

    // 错误编码
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 成功标记
    
    Succeeded   bool `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

}
