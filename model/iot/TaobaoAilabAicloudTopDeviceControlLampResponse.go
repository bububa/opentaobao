package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
台灯控制 APIResponse
taobao.ailab.aicloud.top.device.control.lamp

台灯控制
*/
type TaobaoAilabAicloudTopDeviceControlLampAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ailab_aicloud_top_device_control_lamp_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AiCloudResult `json:"result,omitempty" xml:"