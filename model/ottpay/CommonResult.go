package ottpay

// CommonResult 
type CommonResult struct {

    // data
    
    Data   *TvOrderResultDTO `json:"data,omitempty" xml:"data,omitempty"`
    

    // 返回码
    
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    

    // 错误信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 返回结果
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
