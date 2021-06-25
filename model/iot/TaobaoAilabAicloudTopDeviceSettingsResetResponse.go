package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
重置设备个性化设置 APIResponse
taobao.ailab.aicloud.top.device.settings.reset

重置设备个性化设置
*/
type TaobaoAilabAicloudTopDeviceSettingsResetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAilabAicloudTopDeviceSettingsResetResponse `json:"taobao_ailab_aicloud_top_device_settings_reset_response,omitempty"`
}

type TaobaoAilabAicloudTopDeviceSettingsResetResponse struct {

    // 业务结果是否成功
    Model   bool `json:"model,omitempty"`

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty"`

    // 网络请求是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 扩展字段
    ExtendInfo   string `json:"extend_info,omitempty"`

}