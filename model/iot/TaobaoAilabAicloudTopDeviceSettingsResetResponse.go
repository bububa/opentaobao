package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
重置设备个性化设置 APIResponse
taobao.ailab.aicloud.top.device.settings.reset

重置设备个性化设置
*/
type TaobaoAilabAicloudTopDeviceSettingsResetAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopDeviceSettingsResetResponse
}

type TaobaoAilabAicloudTopDeviceSettingsResetResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_device_settings_reset_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 业务结果是否成功
    
    Model   bool `json:"model,omitempty" xml:"model,omitempty"`

    
    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // 网络请求是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 扩展字段
    
    ExtendInfo   string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`

    
}
