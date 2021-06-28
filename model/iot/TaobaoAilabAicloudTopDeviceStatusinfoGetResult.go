package iot

// TaobaoAilabAicloudTopDeviceStatusinfoGetResult 
/* model for simplify = false
type TaobaoAilabAicloudTopDeviceStatusinfoGetResult struct {

    // code
    
    Code   int64 `json:"code,omitempty"`
    

    // message
    
    Message   string `json:"message,omitempty"`
    

    // 设备状态信息
    
    Result  *struct {
        TopDeviceStatusInfoDto  *TopDeviceStatusInfoDto `json:"top_device_status_info_dto,omitempty"`
    } `json:"result,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

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
