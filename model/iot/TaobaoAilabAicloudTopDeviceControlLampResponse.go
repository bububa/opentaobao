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
    Response *TaobaoAilabAicloudTopDeviceControlLampResponse `json:"taobao_ailab_aicloud_top_device_control_lamp_response,omitempty"`
}

type TaobaoAilabAicloudTopDeviceControlLampResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
