package wdk

// HmResult 
type HmResult struct {

    // 设备列表
    Models   []DeviceInfoDto `json:"models,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // model
    Model   *MqttDeviceInfoDto `json:"model,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

}
