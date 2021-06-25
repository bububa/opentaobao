package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
点播url APIResponse
taobao.ailab.aicloud.top.device.control.playurl

点播url
*/
type TaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAilabAicloudTopDeviceControlPlayurlResponse `json:"taobao_ailab_aicloud_top_device_control_playurl_response,omitempty"`
}

type TaobaoAilabAicloudTopDeviceControlPlayurlResponse struct {

    // 业务请求是否成功
    Model   bool `json:"model,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // 网络请求是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
