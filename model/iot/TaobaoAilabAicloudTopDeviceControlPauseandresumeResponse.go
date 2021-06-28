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
    // Response *TaobaoAilabAicloudTopDeviceControlPauseandresumeResponse `json:"ailab_aicloud_top_device_control_pauseandresume_response,omitempty"` 
    TaobaoAilabAicloudTopDeviceControlPauseandresumeResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopDeviceControlPauseandresumeResponse struct {

    // result
    
    Result  *struct {
        AiCloudResult  *AiCloudResult `json:"ai_cloud_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopDeviceControlPauseandresumeResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
