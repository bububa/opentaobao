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
	RequestId     string         `json:"request_id,omitempty" xml:"ailab_aicloud_top_device_settings_reset_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 业务结果是否成功
    
    Model   bool `json:"model,omitempty" xml:"