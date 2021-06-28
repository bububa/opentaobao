package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
台灯控制 APIResponse
taobao.ailab.aicloud.top.device.control.lamp

台灯控制
*/
type TaobaoAilabAicloudTopDeviceControlLampAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopDeviceControlLampResponse `json:"ailab_aicloud_top_device_control_lamp_response,omitempty"` 
    TaobaoAilabAicloudTopDeviceControlLampResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopDeviceControlLampResponse struct {

    // result
    
    Result  *struct {
        AiCloudResult  *AiCloudResult `json:"ai_cloud_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopDeviceControlLampResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
