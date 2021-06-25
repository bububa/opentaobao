package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
设备控制自定义扩展接口 APIResponse
taobao.ailab.aicloud.top.device.control.custom

设备控制自定义扩展接口
*/
type TaobaoAilabAicloudTopDeviceControlCustomAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAilabAicloudTopDeviceControlCustomResponse `json:"taobao_ailab_aicloud_top_device_control_custom_response,omitempty"`
}

type TaobaoAilabAicloudTopDeviceControlCustomResponse struct {

    // 业务请求是否成功
    Model   bool `json:"model,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // 网络请求是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
