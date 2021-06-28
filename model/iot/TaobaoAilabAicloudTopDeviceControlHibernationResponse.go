package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
定时休眠 APIResponse
taobao.ailab.aicloud.top.device.control.hibernation

定时休眠
*/
type TaobaoAilabAicloudTopDeviceControlHibernationAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAilabAicloudTopDeviceControlHibernationResponse `json:"ailab_aicloud_top_device_control_hibernation_response,omitempty"` 
    TaobaoAilabAicloudTopDeviceControlHibernationResponse
}

/* model for simplify = false
type TaobaoAilabAicloudTopDeviceControlHibernationResponse struct {

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

type TaobaoAilabAicloudTopDeviceControlHibernationResponse struct {

    // 业务请求是否成功
    Model   bool `json:"model,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // 网络请求是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
