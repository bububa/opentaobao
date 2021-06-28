package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取设备授权码 APIResponse
taobao.ailab.aicloud.top.device.authcode.get

获取设备授权码
*/
type TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopDeviceAuthcodeGetResponse `json:"ailab_aicloud_top_device_authcode_get_response,omitempty"` 
    TaobaoAilabAicloudTopDeviceAuthcodeGetResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopDeviceAuthcodeGetResponse struct {

    // 设备授权码，后续流程中所述的auth code
    
    Model   string `json:"model,omitempty"`
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // msgInfo错误码信息，成功返回null
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

}
*/

type TaobaoAilabAicloudTopDeviceAuthcodeGetResponse struct {

    // 设备授权码，后续流程中所述的auth code
    Model   string `json:"model,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // msgInfo错误码信息，成功返回null
    MsgInfo   string `json:"msg_info,omitempty"`

}
