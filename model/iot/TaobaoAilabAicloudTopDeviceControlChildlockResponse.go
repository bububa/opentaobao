package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
设备儿童锁 APIResponse
taobao.ailab.aicloud.top.device.control.childlock

设备儿童锁
*/
type TaobaoAilabAicloudTopDeviceControlChildlockAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopDeviceControlChildlockResponse `json:"ailab_aicloud_top_device_control_childlock_response,omitempty"` 
    TaobaoAilabAicloudTopDeviceControlChildlockResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopDeviceControlChildlockResponse struct {

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

type TaobaoAilabAicloudTopDeviceControlChildlockResponse struct {

    // 业务请求是否成功
    Model   bool `json:"model,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // 网络请求是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
