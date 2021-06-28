package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
设备音量 APIResponse
taobao.ailab.aicloud.top.device.control.volume

设备音量
*/
type TaobaoAilabAicloudTopDeviceControlVolumeAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopDeviceControlVolumeResponse `json:"ailab_aicloud_top_device_control_volume_response,omitempty"` 
    TaobaoAilabAicloudTopDeviceControlVolumeResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopDeviceControlVolumeResponse struct {

    // 业务请求是否成功
    
    Model   bool `json:"model,omitempty"`
    

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // 网络请求是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopDeviceControlVolumeResponse struct {

    // 业务请求是否成功
    Model   bool `json:"model,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // 网络请求是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
