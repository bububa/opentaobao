package tmallgenie

// Payload 
type Payload struct {

    // 错误码，出错时返回
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 错误信息，出错时返回
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 设备id
    
    DeviceId   string `json:"device_id,omitempty" xml:"device_id,omitempty"`
    

}
