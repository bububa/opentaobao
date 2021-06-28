package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
设备播放暂停 APIResponse
taobao.ailab.aicloud.top.device.control.pauseandresume

设备播放暂停
*/
type TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopDeviceControlPauseandresumeResponse
}

type TaobaoAilabAicloudTopDeviceControlPauseandresumeResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_device_control_pauseandresume_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
