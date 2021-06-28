package wdk

// TopBaseResult 
type TopBaseResult struct {

    // 表示调用是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 返回码说明
    
    ReturnMsg   string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
    

    // 返回码
    
    ReturnCode   string `json:"return_code,omitempty" xml:"return_code,omitempty"`
    

    // 调用处理结果
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`
    

}
