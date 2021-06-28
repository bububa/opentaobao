package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
开放设备id转换内部设备id APIResponse
taobao.ailab.aicloud.top.device.deviceid.convert

将开放设备id转换为内部设备id
*/
type TaobaoAilabAicloudTopDeviceDeviceidConvertAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopDeviceDeviceidConvertResponse `json:"ailab_aicloud_top_device_deviceid_convert_response,omitempty"` 
    TaobaoAilabAicloudTopDeviceDeviceidConvertResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopDeviceDeviceidConvertResponse struct {

    // 接口返回model
    
    Result  *struct {
        TaobaoAilabAicloudTopDeviceDeviceidConvertResult  *TaobaoAilabAicloudTopDeviceDeviceidConvertResult `json:"taobao_ailab_aicloud_top_device_deviceid_convert_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopDeviceDeviceidConvertResponse struct {

    // 接口返回model
    Result   *TaobaoAilabAicloudTopDeviceDeviceidConvertResult `json:"result,omitempty"`

}
