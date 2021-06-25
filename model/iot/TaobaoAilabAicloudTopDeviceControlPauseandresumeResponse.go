package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
设备播放暂停 APIResponse
taobao.ailab.aicloud.top.device.control.pauseandresume

设备播放暂停
*/
type TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAilabAicloudTopDeviceControlPauseandresumeResponse `json:"taobao_ailab_aicloud_top_device_control_pauseandresume_response,omitempty"`
}

type TaobaoAilabAicloudTopDeviceControlPauseandresumeResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
