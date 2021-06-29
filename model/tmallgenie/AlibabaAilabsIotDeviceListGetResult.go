package tmallgenie

// AlibabaAilabsIotDeviceListGetResult 
type AlibabaAilabsIotDeviceListGetResult struct {

    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 错误码
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
    

    // 返回值list
    
    RetValues   []RetValue `json:"ret_values,omitempty" xml:"ret_values,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
