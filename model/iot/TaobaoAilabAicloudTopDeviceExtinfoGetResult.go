package iot

// TaobaoAilabAicloudTopDeviceExtinfoGetResult 
type TaobaoAilabAicloudTopDeviceExtinfoGetResult struct {

    // message
    Message   string `json:"message,omitempty"`

    // 三方设备信息
    Result   *TopDeviceExtInfoDto `json:"result,omitempty"`

    // code
    Code   int64 `json:"code,omitempty"`

}
