package iot

// TaobaoAilabAicloudTopDeviceDetailinfoGetResult 
/* model for simplify = false
type TaobaoAilabAicloudTopDeviceDetailinfoGetResult struct {

    // code
    
    Code   int64 `json:"code,omitempty"`
    

    // message
    
    Message   string `json:"message,omitempty"`
    

    // 设备详细信息
    
    Result  *struct {
        TopDeviceDetailInfoDto  *TopDeviceDetailInfoDto `json:"top_device_detail_info_dto,omitempty"`
    } `json:"result,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TaobaoAilabAicloudTopDeviceDetailinfoGetResult 
type TaobaoAilabAicloudTopDeviceDetailinfoGetResult struct {

    // code
    Code   int64 `json:"code,omitempty"`

    // message
    Message   string `json:"message,omitempty"`

    // 设备详细信息
    Result   *TopDeviceDetailInfoDto `json:"result,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}
