package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
通过id播放歌曲 APIResponse
taobao.ailab.aicloud.top.device.control.playbyid

通过id播放歌曲
*/
type TaobaoAilabAicloudTopDeviceControlPlaybyidAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopDeviceControlPlaybyidResponse `json:"ailab_aicloud_top_device_control_playbyid_response,omitempty"` 
    TaobaoAilabAicloudTopDeviceControlPlaybyidResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopDeviceControlPlaybyidResponse struct {

    // 返回结果
    
    Result  *struct {
        AiCloudResult  *AiCloudResult `json:"ai_cloud_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopDeviceControlPlaybyidResponse struct {

    // 返回结果
    Result   *AiCloudResult `json:"result,omitempty"`

}
