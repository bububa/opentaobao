package iot

// TaobaoAilabAicloudTopDeviceExtinfoGetResult 
/* model for simplify = false
type TaobaoAilabAicloudTopDeviceExtinfoGetResult struct {

    // message
    
    Message   string `json:"message,omitempty"`
    

    // 三方设备信息
    
    Result  *struct {
        TopDeviceExtInfoDto  *TopDeviceExtInfoDto `json:"top_device_ext_info_dto,omitempty"`
    } `json:"result,omitempty"`
    

    // code
    
    Code   int64 `json:"code,omitempty"`
    

}
*/

// TaobaoAilabAicloudTopDeviceExtinfoGetResult 
type TaobaoAilabAicloudTopDeviceExtinfoGetResult struct {

    // message
    Message   string `json:"message,omitempty"`

    // 三方设备信息
    Result   *TopDeviceExtInfoDto `json:"result,omitempty"`

    // code
    Code   int64 `json:"code,omitempty"`

}
