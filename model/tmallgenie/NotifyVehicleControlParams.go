package tmallgenie

// NotifyVehicleControlParams 
type NotifyVehicleControlParams struct {

    // 标准控制协议中的payload
    
    Payload   *Payload `json:"payload,omitempty" xml:"payload,omitempty"`
    

    // 标准控制协议中的header
    
    Header   *IotCommonHeader `json:"header,omitempty" xml:"header,omitempty"`
    

    // 标准查询协议中的exceptions，异常检测项，如果有则返回，没有则不返回
    
    Exceptions   []IotCommonDeviceProperty `json:"exceptions,omitempty" xml:"exceptions,omitempty"`
    

    // 标准查询协议中的properties，异常检测项，如果有则返回，没有则不返回
    
    Properties   []IotCommonDeviceProperty `json:"properties,omitempty" xml:"properties,omitempty"`
    

}
