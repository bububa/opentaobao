package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
解绑设备 APIResponse
taobao.ailab.aicloud.top.device.unbind

解绑设备
*/
type TaobaoAilabAicloudTopDeviceUnbindAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopDeviceUnbindResponse `json:"ailab_aicloud_top_device_unbind_response,omitempty"` 
    TaobaoAilabAicloudTopDeviceUnbindResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopDeviceUnbindResponse struct {

    // 解绑是否成功
    
    Model   bool `json:"model,omitempty"`
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // msgInfo错误码信息，成功返回null
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopDeviceUnbindResponse struct {

    // 解绑是否成功
    Model   bool `json:"model,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // msgInfo错误码信息，成功返回null
    MsgInfo   string `json:"msg_info,omitempty"`

}
