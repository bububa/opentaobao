package iot

// TaobaoAilabAicloudTopDeviceDeviceidConvertResult 
/* model for simplify = false
type TaobaoAilabAicloudTopDeviceDeviceidConvertResult struct {

    // 返回成功
    
    Code   int64 `json:"code,omitempty"`
    

    // 描述
    
    Message   string `json:"message,omitempty"`
    

    // 结果详情
    
    Result  *struct {
        TopDeviceBaseInfoDto  *TopDeviceBaseInfoDto `json:"top_device_base_info_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

// TaobaoAilabAicloudTopDeviceDeviceidConvertResult 
type TaobaoAilabAicloudTopDeviceDeviceidConvertResult struct {

    // 返回成功
    Code   int64 `json:"code,omitempty"`

    // 描述
    Message   string `json:"message,omitempty"`

    // 结果详情
    Result   *TopDeviceBaseInfoDto `json:"result,omitempty"`

}
