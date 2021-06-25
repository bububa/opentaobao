package iot

// TaobaoAilabAicloudTopDeviceStatusinfoGetResult 
type TaobaoAilabAicloudTopDeviceStatusinfoGetResult struct {

    // code
    Code   int64 `json:"code,omitempty"`

    // message
    Message   string `json:"message,omitempty"`

    // 设备状态信息
    Result   *TopDeviceStatusInfoDto `json:"result,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}
